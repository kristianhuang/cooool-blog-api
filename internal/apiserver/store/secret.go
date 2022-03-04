/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package store

import (
	"context"

	"blog-api/internal/pkg/model"
	metav1 "blog-api/pkg/meta/v1"
)

type SecretStore interface {
	Create(ctx context.Context, secret *model.Secret, opts metav1.CreateOptions) error
	Update(ctx context.Context, secret *model.Secret, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username, name string, opts metav1.GetOptions) (*model.Secret, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*model.SecretList, error)
}
