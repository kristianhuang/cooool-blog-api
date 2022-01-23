/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"blog-api/internal/apiserver/service/v1/adminuser"
	"blog-api/internal/apiserver/store"
)

type Service interface {
	AdminUser() adminuser.AdminUserService
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{store: store}
}

func (s *service) AdminUser() adminuser.AdminUserService {
	return adminuser.NewAdminUserService(s.store)
}
