/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"blog-api/internal/pkg/response"
	metav1 "blog-api/pkg/meta/v1"
	"github.com/gin-gonic/gin"
)

func (a *AdminUserController) DeleteCollection(c *gin.Context) {
	if err := a.srv.AdminUser().DeleteCollection(c, c.QueryArray("name"), metav1.DeleteOptions{}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, nil)
}
