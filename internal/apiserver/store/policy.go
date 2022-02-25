/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package store

import (
	"context"

	v1 "blog-api/internal/pkg/model"
	metav1 "blog-api/pkg/meta/v1"
)

type PolicyStore interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
}
