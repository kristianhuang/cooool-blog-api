/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"context"

	modelv1 "blog-api/internal/pkg/model/v1"
	metav1 "blog-api/pkg/meta/v1"
	"gorm.io/gorm"
)

type adminUser struct {
	db *gorm.DB
}

func newAdminUser(db *gorm.DB) *adminUser {
	return &adminUser{db: db}
}

func (u *adminUser) Create(ctx context.Context, adminUserModel *modelv1.AdminUser, opts metav1.CreateOptions) error {
	return u.db.Create(adminUserModel).Error
}

// func (u adminUser) List(cxt context.Context, opts metav1.ListOptions) (*modelv1.AdminUserList, error) {
// 	userList := &modelv1.AdminUserList{}
// 	ol := gormutil.Unpointer(opts.Offset, opts.Limit)
//
// 	where := &modelv1.AdminUser{}
//
// }
