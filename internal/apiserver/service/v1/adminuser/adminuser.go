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
	Update(ctx context.Context, user *model.AdminUser, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, account string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, accounts []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, account string, opts metav1.GetOptions) (*model.AdminUser, error)
	List(ctx context.Context, opts metav1.ListOptions) (*model.AdminUserList, error)
	ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*model.AdminUserList, error)
	ChangePassword(ctx context.Context, user *model.AdminUser) error
}

type adminUserService struct {
	store store.Factory
}

func NewAdminUserService(s store.Factory) *adminUserService {
	return &adminUserService{store: s}
}
