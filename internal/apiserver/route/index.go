/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"blog-go/internal/apiserver/controller/v1/index"
	"github.com/gin-gonic/gin"
)

func IndexRoutes(e *gin.Engine) {
	router := index.NewIndexController()
	e.GET("/", router.Index)
}
