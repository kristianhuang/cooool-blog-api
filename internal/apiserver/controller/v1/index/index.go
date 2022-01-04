/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (ic *IndexController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "index",
	})
}

func NewIndexController() *IndexController {
	return &IndexController{}
}
