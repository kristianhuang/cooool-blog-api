/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"cooool-blog-api/internal/apiserver/controller/v1/adminuser"

	"github.com/gin-gonic/gin"
)

func AdminUser(e *gin.Engine) {
	v1 := e.Group(V1)
	{
		adminUserController := adminuser.NewController()
		adminUserV1 := v1.Group("/admin-user")
		{
			adminUserV1.POST("", adminUserController.Create)
			adminUserV1.DELETE("", adminUserController.DeleteCollection)
			adminUserV1.DELETE(":username", adminUserController.Delete)
			adminUserV1.PUT(":username", adminUserController.Update)
			adminUserV1.PUT(":username/change-password", adminUserController.ChangePassword)
			adminUserV1.GET("", adminUserController.List)
			adminUserV1.GET(":username", adminUserController.Get)
		}
	}
}
