/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"context"

	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/model"
	"blog-api/internal/pkg/util/gormutil"
	"blog-api/pkg/errors"
	"blog-api/pkg/fields"
	metav1 "blog-api/pkg/meta/v1"
	"gorm.io/gorm"
)

type policies struct {
	db *gorm.DB
}

func newPolicy(ds *dataStore) *policies {
	return &policies{ds.db}
}

func (p *policies) Create(ctx context.Context, policyModel *model.Policy, opts metav1.CreateOptions) error {
	return p.db.Create(&policyModel).Error
}

func (p *policies) Update(ctx context.Context, policyModel *model.Policy, opts metav1.UpdateOptions) error {
	return p.db.Save(&policyModel).Error
}

func (p *policies) Delete(ctx context.Context, account, name string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	err := p.db.Where("account = ? AND name = ?", account, name).Delete(&model.Policy{}).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p policies) DeleteByName(cxt context.Context, name string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("name = ?", name).Delete(&model.Policy{}).Error
}

func (p *policies) DeleteCollection(ctx context.Context, name string, accounts []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("name = ? AND account IN (?)", name, accounts).Delete(&model.Policy{}).Error
}

func (p *policies) Get(ctx context.Context, name, account string, opts metav1.GetOptions) (*model.Policy, error) {
	policy := &model.Policy{}
	if err := p.db.Where("name = ? AND account = ?", name, account).First(policy).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrPolicyNotFound, err.Error())
		}
	}

	return policy, nil
}

func (p *policies) List(ctx context.Context, account string, opts metav1.ListOptions) (*model.PolicyList, error) {
	ret := &model.PolicyList{}
	ol := gormutil.Unpointer(opts.Limit, opts.Offset)

	if account != "" {
		p.db = p.db.Where("account = ?", account)
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	name, _ := selector.RequiresExactMatch("name")

	d := p.db.Where("name like ?", "%"+name+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
