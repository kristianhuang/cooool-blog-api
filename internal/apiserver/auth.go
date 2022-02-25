/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	"blog-api/internal/apiserver/store"
	"blog-api/internal/pkg/middleware"
	"blog-api/internal/pkg/middleware/auth"
	metav1 "blog-api/pkg/meta/v1"
	log "blog-api/pkg/rollinglog"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "blog-admin"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "blog-apiserver"
)

type loginInfo struct {
	Account  string `form:"account" json:"account" binding:"required,account"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func newBasicAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(account, password string) bool {
		adminUser, err := store.Client().AdminUser().Get(context.TODO(), account, metav1.GetOptions{})
		if err != nil {
			return false
		}

		if err := adminUser.Compare(password); err != nil {
			return false
		}

		adminUser.LoginedAt = time.Now().Unix()
		_ = store.Client().AdminUser().Update(context.TODO(), adminUser, metav1.UpdateOptions{})

		return true
	})
}

// TODO JWT Server
func newJWTAuth() {

}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var (
			login loginInfo
			err   error
		)

		if c.Request.Header.Get("Authorization") != "" {
			login, err = parseWithHeader(c)
		} else {
			login, err = parseWithBody(c)
		}

		if err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		adminUser, err := store.Client().AdminUser().Get(c, login.Account, metav1.GetOptions{})

		if err != nil {
			log.Errorf("get user information failed: %s", err.Error())

			return "", jwt.ErrFailedAuthentication
		}

		if err := adminUser.Compare(login.Password); err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		adminUser.LoginedAt = time.Now().Unix()
		_ = store.Client().AdminUser().Update(context.TODO(), adminUser, metav1.UpdateOptions{})

		return adminUser, nil
	}
}

func parseWithHeader(c *gin.Context) (loginInfo, error) {
	authz := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(authz) != 2 || authz[0] != "Basic" {
		log.Errorf("get basic string from Authorization header failed")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	payload, err := base64.StdEncoding.DecodeString(authz[1])
	if err != nil {
		log.Errorf("decode basic string: %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		log.Errorf("parse payload failed")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return loginInfo{
		Account:  pair[0],
		Password: pair[1],
	}, nil
}

func parseWithBody(c *gin.Context) (loginInfo, error) {
	var login loginInfo
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Errorf("parse login parameters: %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return login, nil
}
