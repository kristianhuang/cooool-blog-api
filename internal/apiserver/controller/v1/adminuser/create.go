/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"net/http"

	"blog-api/internal/pkg/model"
	metav1 "blog-api/pkg/meta/v1"
	log "blog-api/pkg/rollinglog"
	"github.com/gin-gonic/gin"
)

func (a *AdminUserController) Create(c *gin.Context) {
	log.L(c).Info("user create function called.")

	var adminUser model.AdminUser
	if err := c.ShouldBindJSON(&adminUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg2": err.Error()})
		return
	}

	if err := a.srv.AdminUser().Create(c, &adminUser, metav1.CreateOptions{}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
