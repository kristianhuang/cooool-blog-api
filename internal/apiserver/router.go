/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"cooool-blog-api/internal/apiserver/route"
	"cooool-blog-api/internal/pkg/middleware/auth"

	"github.com/gin-gonic/gin"
)

type Route func(e *gin.Engine)

var (
	Routes = []Route{
		route.Login,
		route.AdminUser,
		route.Secrets,
		route.Policies,
		route.Error,
	}
)

func initRouter(e *gin.Engine) *gin.Engine {
	route.WithAuth(newJWTAuth().(auth.JWTStrategy), newAutoAuth())

	for _, r := range Routes {
		r(e)
	}

	return e
}
