/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"sync"

	"blog-api/internal/apiserver"
	"blog-api/internal/pkg/middleware/auth"
)

const (
	V1 = "/v1"
	// If you need more version...
)

var (
	once        sync.Once
	jwtStrategy auth.JWTStrategy
)

func init() {
	once.Do(func() {
		jwtStrategy, _ = apiserver.NewJWTAuth().(auth.JWTStrategy)
	})
}
