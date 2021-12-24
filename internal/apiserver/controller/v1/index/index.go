/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package index

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Controller struct {
}

func (ic *Controller) Index(c *gin.Context) {
	list := viper.AllSettings()
	fmt.Println(1233)
	c.JSON(http.StatusOK, gin.H{"list": list})
}

func NewIndexController() *Controller {
	return &Controller{}
}
