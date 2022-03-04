/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"blog-api/internal/authzserver/controller/v1/authorize"
	"github.com/gin-gonic/gin"
)

func Authz(e *gin.Engine) {
	v1 := e.Group(V1, authStrategy.AuthFunc())
	{
		authzController := authorize.NewAuthzController(cacheIns)
		v1.POST("/authz", authzController.Authorize)
	}
}
