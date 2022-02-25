/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package store

import (
	"context"

	v1 "blog-api/internal/pkg/model"
	metav1 "blog-api/pkg/meta/v1"
)

type AdminUserStore interface {
	Create(ctx context.Context, adminUserModel *v1.AdminUser, opts metav1.CreateOptions) error

	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.AdminUser, error)

	Update(ctx context.Context, adminUserModel *v1.AdminUser, opts metav1.UpdateOptions) error

	Delete(ctx context.Context, account string, opts metav1.DeleteOptions) error
}
