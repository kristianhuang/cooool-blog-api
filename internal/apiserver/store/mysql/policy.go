/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"context"

	"blog-api/internal/pkg/model"
	metav1 "blog-api/pkg/meta/v1"
	"gorm.io/gorm"
)

type policy struct {
	db *gorm.DB
}

func newPolicy(ds *dataStore) *policy {
	return &policy{ds.db}
}

func (p *policy) Create(ctx context.Context, policyModel *model.Policy, opts metav1.CreateOptions) error {

	return p.db.Create(&policyModel).Error
}
