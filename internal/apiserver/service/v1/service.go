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
	Policy() PolicyService
	Secrets() SecretService
}

type service struct {
	store store.Factory
}

func NewService() Service {
	return &service{
		store: store.Client(),
	}
}

func (s *service) AdminUser() AdminUserService {
	return newAdminUserService(s.store)
}

func (s *service) Policy() PolicyService {
	return newPolicies(s.store)
}

func (s *service) Secrets() SecretService {
	return newSecrets(s)
}
