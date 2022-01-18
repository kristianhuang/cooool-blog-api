/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package middleware

import "github.com/gin-gonic/gin"

var Middlewares = registerMiddlewares()

// Register global middleware.
func registerMiddlewares() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"recovery":  gin.Recovery(),
		"cors":      Cors(),
		"requestid": RequestID(),
	}
}
