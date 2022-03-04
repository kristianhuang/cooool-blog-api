/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package secret

import (
	srvv1 "blog-api/internal/apiserver/service/v1"
	"blog-api/internal/pkg/bind"
)

type SecretController struct {
	srv srvv1.Service
	*bind.BindData
}

func NewSecretController() *SecretController {
	return &SecretController{
		srv:      srvv1.NewService(),
		BindData: bind.NewBindData(),
	}
}
