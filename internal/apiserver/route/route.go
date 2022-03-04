/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"blog-api/internal/pkg/middleware"
	middlewareauth "blog-api/internal/pkg/middleware/auth"
)

const (
	V1 = "/v1"
	// If you need more version...
)

type auth struct {
	jwt  middlewareauth.JWTStrategy
	auto middleware.AuthStrategy
}

var (
	authStrategy auth
)

func WithAuth(jwt middlewareauth.JWTStrategy, auto middleware.AuthStrategy) {
	authStrategy = auth{
		jwt:  jwt,
		auto: auto,
	}
}
