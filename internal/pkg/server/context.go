/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import "github.com/gin-gonic/gin"

const UseKey = "clientUser"

// TODO 记录当前请求客户 key 的包
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
