/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"cooool-blog-api/internal/authzserver/load/cache"
	"cooool-blog-api/internal/pkg/middleware"
)

const (
	V1 = "/v1"
	// If you need more version...
)

var (
	authStrategy middleware.AuthStrategy
	cacheIns     *cache.Cache
)

func WithAuth(auto middleware.AuthStrategy) {
	authStrategy = auto
}

func WithCacheIns(c *cache.Cache) {
	cacheIns = c
}
