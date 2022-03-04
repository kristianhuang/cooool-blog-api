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

func (s *SecretController) List(c *gin.Context) {
	var r metav1.ListOptions
	if err := s.Bind(c, &r); err != nil {
		response.Write(c, err, nil)
		return
	}

	secrets, err := s.srv.Secrets().List(
		c,
		c.GetString(middleware.UsernameKey),
		r,
	)
	if err != nil {
		response.Write(c, err, nil)

		return
	}

	response.Write(c, nil, secrets)
}
