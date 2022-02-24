/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package middleware

import (
	log "blog-api/pkg/rollinglog"
	"github.com/gin-gonic/gin"
)

// UsernameKey defines the key in gin context which represents the owner of the secret.
const UsernameKey = "username"

// Context is a middleware that injects common prefix fields to gin.Context.
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(log.KeyRequestID, c.GetString(XRequestIDKey))
		c.Set(log.KeyUsername, c.GetString(UsernameKey))
		c.Next()
	}
}
