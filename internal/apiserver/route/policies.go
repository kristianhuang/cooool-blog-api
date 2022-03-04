/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"blog-api/internal/apiserver/controller/v1/policy"
	"github.com/gin-gonic/gin"
)

func Policies(e *gin.Engine) {
	v1 := e.Group(V1)
	{
		policyController := policy.NewPolicyController()
		policyV1 := v1.Group("/policies")
		policyV1.Use(authStrategy.auto.AuthFunc())
		{
			policyV1.POST("", policyController.Create)
			policyV1.DELETE(":name", policyController.Delete)
			policyV1.PUT(":name", policyController.Delete)
			policyV1.GET(":name", policyController.Get)
			policyV1.GET("", policyController.List)
		}
	}

}
