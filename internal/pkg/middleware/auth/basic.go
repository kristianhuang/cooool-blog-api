/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package auth

import (
	"encoding/base64"
	"strings"

	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/internal/pkg/middleware"
	"cooool-blog-api/internal/pkg/response"
	"cooool-blog-api/pkg/errors"

	"github.com/gin-gonic/gin"
)

type BasicStrategy struct {
	compare func(username, password string) bool
}

var _ middleware.AuthStrategy = &BasicStrategy{}

func NewBasicStrategy(compare func(username, password string) bool) BasicStrategy {
	return BasicStrategy{compare: compare}
}

func (s BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			response.Write(c, errors.WithCode(code.ErrInvalidAuthHeader, "Authorization header format is wrong."), nil)

			c.Abort()
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || !s.compare(pair[0], pair[1]) {
			response.Write(c, errors.WithCode(code.ErrInvalidAuthHeader, "Authorization header format is wrong."), nil)

			c.Abort()
			return
		}

		c.Set(middleware.UsernameKey, pair[0])

		c.Next()
	}
}
