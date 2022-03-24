/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"cooool-blog-api/internal/apiserver/controller/v1/secret"

	"github.com/gin-gonic/gin"
)

func Secrets(e *gin.Engine) {
	v1 := e.Group(V1)
	{
		secretController := secret.NewSecretController()
		secretV1 := v1.Group("/secrets")
		secretV1.Use(authStrategy.auto.AuthFunc())
		{
			secretV1.POST("", secretController.Create)
			secretV1.DELETE(":name", secretController.Delete)
			secretV1.PUT(":name", secretController.Delete)
			secretV1.GET(":name", secretController.Get)
			secretV1.GET("", secretController.List)
		}
	}
}
