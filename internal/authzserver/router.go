/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authzserver

import (
	"blog-api/internal/authzserver/load/cache"
	"blog-api/internal/authzserver/route"
	log "blog-api/pkg/rollinglog"
	"github.com/gin-gonic/gin"
)

type Route func(e *gin.Engine)

var (
	Routes = []Route{
		route.Authz,
		route.Error,
	}
)

func initRouter(e *gin.Engine) *gin.Engine {
	cacheIns, _ := cache.GetCacheInsOr(nil)
	if cacheIns == nil {
		log.Panicf("get nil cache instance")
	}
	route.WithCacheIns(cacheIns)
	route.WithAuth(newCacheAuth())

	for _, r := range Routes {
		r(e)
	}

	return e
}
