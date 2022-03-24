/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package policy

import (
	srvv1 "cooool-blog-api/internal/apiserver/service/v1"
	"cooool-blog-api/internal/pkg/bind"
)

type PolicyController struct {
	srv srvv1.Service
	*bind.BindData
}

func NewPolicyController() *PolicyController {
	return &PolicyController{
		srv:      srvv1.NewService(),
		BindData: bind.NewBindData(),
	}
}
