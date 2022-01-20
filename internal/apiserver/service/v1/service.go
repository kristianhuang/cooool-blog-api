/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"blog-api/internal/apiserver/store"
)

type Service interface {
	AdminUser() AdminUserService
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) *service {
	return &service{store: store}
}

func (s *service) AdminUser() AdminUserService {
	return newAdminUserService(s)
}
