/*
 * Copyright 2021 KristianHuang <KristianHuangyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	srvv1 "blog-api/internal/apiserver/service/v1"
	"blog-api/internal/apiserver/store"
)

type AdminUserController struct {
	srv srvv1.Service
}

func NewController(store store.Factory) *AdminUserController {
	return &AdminUserController{
		srv: srvv1.NewService(store),
	}
}
