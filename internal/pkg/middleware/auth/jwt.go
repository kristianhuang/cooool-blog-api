/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package auth

import (
	"cooool-blog-api/internal/pkg/middleware"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// AuthAudience defines the value of jwt audience field.
const AuthAudience = "blog-admin"

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ middleware.AuthStrategy = &JWTStrategy{}

func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{GinJWTMiddleware: gjwt}
}

func (s JWTStrategy) AuthFunc() gin.HandlerFunc {
	return s.MiddlewareFunc()
}
