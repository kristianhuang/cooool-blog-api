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

func (p *policies) Delete(ctx context.Context, name, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	err := p.db.Where("username = ? and name = ?", username, name).Delete(&model.Policy{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (p policies) DeleteByAdminUser(cxt context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username = ?", username).Delete(&model.Policy{}).Error
}

func (p *policies) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username = ? AND name IN (?)", username, names).Delete(&model.Policy{}).Error
}

func (p policies) DeleteCollectionByAdminUser(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("user_name IN (?)", usernames).Delete(&model.Policy{}).Error
}

func (p *policies) Get(ctx context.Context, username, name string, opts metav1.GetOptions) (*model.Policy, error) {
	policy := &model.Policy{}

	err := p.db.Where("username = ? AND name = ?", username, username).First(policy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return policy, nil
}

func (p *policies) List(ctx context.Context, username string, opts metav1.ListOptions) (*model.PolicyList, error) {
	ret := &model.PolicyList{}
	ol := gormutil.Unpointer(opts.Limit, opts.Offset)

	if username != "" {
		p.db = p.db.Where("user_name = ?", username)
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	name, _ := selector.RequiresExactMatch("name")

	d := p.db.Where("username like ?", "%"+name+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
