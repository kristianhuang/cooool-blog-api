/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package auth

import (
	"fmt"
	"time"

	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/internal/pkg/middleware"
	"cooool-blog-api/internal/pkg/response"
	"cooool-blog-api/pkg/errors"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

var (
	ErrMissingKID    = errors.New("Invalid token format: missing kid field in claims")
	ErrMissingSecret = errors.New("Can not obtain secret information from cache")
)

type Secret struct {
	Username string
	ID       string
	Key      string
	Expires  int64
}

type CacheStrategy struct {
	get func(kid string) (Secret, error)
}

var _ middleware.AuthStrategy = &CacheStrategy{}

func NewCacheStrategy(get func(kid string) (Secret, error)) CacheStrategy {
	return CacheStrategy{get: get}
}

func (s CacheStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if len(header) == 0 {
			response.Write(c, errors.WithCode(code.ErrMissingHeader, "Authorization header cannot be empty."), nil)
			c.Abort()

			return
		}
		var rawJWT string
		// Parse the header to get the token part.
		fmt.Sscanf(header, "Bearer %s", &rawJWT)

		// Use own validation logic, see below
		var secret Secret

		claims := &jwt.MapClaims{}
		// Verify the token
		parsedT, err := jwt.ParseWithClaims(rawJWT, claims, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is HMAC signature
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, ErrMissingKID
			}

			var err error
			secret, err = s.get(kid)
			if err != nil {
				return nil, ErrMissingSecret
			}

			return []byte(secret.Key), nil
		}, jwt.WithAudience(AuthAudience))
		if err != nil || !parsedT.Valid {
			response.Write(c, errors.WithCode(code.ErrSignatureInvalid, err.Error()), nil)
			c.Abort()

			return
		}

		if KeyExpired(secret.Expires) {
			tm := time.Unix(secret.Expires, 0).Format("2006-01-02 15:04:05")
			response.Write(c, errors.WithCode(code.ErrExpired, "expired at: %s", tm), nil)
			c.Abort()

			return
		}

		c.Set(middleware.UsernameKey, secret.Username)
		c.Next()
	}
}

// KeyExpired checks if a key has expired, if the value of user.SessionState.Expires is 0, it will be ignored.
func KeyExpired(expires int64) bool {
	if expires >= 1 {
		return time.Now().After(time.Unix(expires, 0))
	}

	return false
}
