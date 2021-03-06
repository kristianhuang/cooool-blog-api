/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package store

import (
	"context"

	v1 "cooool-blog-api/internal/pkg/model"
	metav1 "cooool-blog-api/pkg/meta/v1"
)

type PolicyStore interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
	Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username, name string, opts metav1.GetOptions) (*v1.Policy, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}
