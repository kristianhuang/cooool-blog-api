/*
 * Copyright 2021 KristianHuang <KristianHuangyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	srvv1 "cooool-blog-api/internal/apiserver/service/v1"
	"cooool-blog-api/internal/pkg/bind"
)

type AdminUserController struct {
	srv srvv1.Service
	*bind.BindData
}

func NewController() *AdminUserController {
	return &AdminUserController{
		srv:      srvv1.NewService(),
		BindData: bind.NewBindData(),
	}
}
