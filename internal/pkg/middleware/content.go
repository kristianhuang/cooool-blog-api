/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package middleware

import "github.com/gin-gonic/gin"

func Content() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
