/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package secret

import (
	"blog-api/internal/pkg/middleware"
	"blog-api/internal/pkg/response"
	metav1 "blog-api/pkg/meta/v1"
	"github.com/gin-gonic/gin"
)

func (s *SecretController) Get(c *gin.Context) {

	secret, err := s.srv.Secrets().Get(
		c,
		c.GetString(middleware.UsernameKey),
		c.Param("name"),
		metav1.GetOptions{},
	)
	if err != nil {
		response.Write(c, err, nil)

		return
	}

	response.Write(c, nil, secret)
}
