/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package policy

import (
	"cooool-blog-api/internal/pkg/middleware"
	"cooool-blog-api/internal/pkg/response"
	metav1 "cooool-blog-api/pkg/meta/v1"

	"github.com/gin-gonic/gin"
)

func (p *PolicyController) Delete(c *gin.Context) {
	if err := p.srv.Policy().Delete(
		c,
		c.GetString(middleware.UsernameKey),
		c.Param("name"),
		metav1.DeleteOptions{},
	); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, nil)
}
