/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-go/internal/apiserver/route"
	"github.com/gin-gonic/gin"
)

type Route func(engine *gin.Engine)

var (
	Routes = []Route{
		route.Index,
	}
)

func initRouter(eg *gin.Engine) *gin.Engine {
	for _, r := range Routes {
		r(eg)
	}

	return eg
}
