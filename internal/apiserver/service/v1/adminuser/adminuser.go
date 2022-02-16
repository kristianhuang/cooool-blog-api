/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"context"

	"blog-api/internal/apiserver/store"
	"blog-api/internal/pkg/model"
	metav1 "blog-api/pkg/meta/v1"
)

type AdminUserService interface {
	Create(ctx context.Context, au *model.AdminUser, options metav1.CreateOptions) error
}

type adminUserService struct {
	store store.Factory
}

func NewAdminUserService(s store.Factory) *adminUserService {
	return &adminUserService{store: s}
}

func (a *adminUserService) Create(ctx context.Context, au *model.AdminUser, options metav1.CreateOptions) error {

	if err := a.store.AdminUser().Create(ctx, au, options); err != nil {
		return err
	}

	return nil
}
