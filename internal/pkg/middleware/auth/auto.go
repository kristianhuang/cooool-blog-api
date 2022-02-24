/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package auth

import (
	"strings"

	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/middleware"
	"blog-api/internal/pkg/response"
	"blog-api/pkg/errors"
	"github.com/gin-gonic/gin"
)

const authHeaderCount = 2

type AutoStrategy struct {
	basic middleware.AuthStrategy
	jwt   middleware.AuthStrategy
}

var _ middleware.AuthStrategy = &AutoStrategy{}

func NewAutoStrategy(basic middleware.AuthStrategy, jwt middleware.AuthStrategy) AutoStrategy {
	return AutoStrategy{basic: basic, jwt: jwt}
}

func (s AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := middleware.AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(authHeader) != authHeaderCount {
			response.Write(c, errors.WithCode(code.ErrInvalidAuthHeader, "Authorization header format is wrong."), nil)

			c.Abort()
			return
		}

		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(s.basic)
		case "Bearer":
			operator.SetStrategy(s.jwt)
			// a.JWT.MiddlewareFunc()(c)
		default:
			response.Write(c, errors.WithCode(code.ErrSignatureInvalid, "unrecognized Authorization header."), nil)
			c.Abort()
			return
		}

		operator.AuthFunc()
		c.Next()
	}
}
