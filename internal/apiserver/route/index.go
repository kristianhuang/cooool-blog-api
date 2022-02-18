/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"blog-api/internal/apiserver/controller/v1/index"
	"github.com/gin-gonic/gin"
)

func Index(e *gin.Engine) {
	indexController := index.NewIndexController()
	e.GET("/", indexController.Index)
}
