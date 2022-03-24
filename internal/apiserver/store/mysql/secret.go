/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"context"

	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/internal/pkg/model"
	"cooool-blog-api/internal/pkg/util/gormutil"
	"cooool-blog-api/pkg/errors"
	"cooool-blog-api/pkg/fields"
	metav1 "cooool-blog-api/pkg/meta/v1"

	"gorm.io/gorm"
)

type secrets struct {
	db *gorm.DB
}

func (s *secrets) Create(ctx context.Context, secret *model.Secret, opts metav1.CreateOptions) error {
	return s.db.Create(secret).Error
}

func (s *secrets) Update(ctx context.Context, secret *model.Secret, opts metav1.UpdateOptions) error {
	return s.db.Save(secret).Error
}

func (s *secrets) Delete(ctx context.Context, username, name string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		s.db = s.db.Unscoped()
	}

	err := s.db.Where("username = ? and name = ?", username, name).Delete(&model.Secret{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (s *secrets) DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		s.db = s.db.Unscoped()
	}

	return s.db.Where("username = ? and name in (?)", username, names).Delete(&model.Secret{}).Error
}

func (s *secrets) Get(ctx context.Context, username, name string, opts metav1.GetOptions) (*model.Secret, error) {
	secret := &model.Secret{}
	err := s.db.Where("username = ? and name= ?", username, name).First(&secret).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrSecretNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return secret, nil
}

func (s *secrets) List(ctx context.Context, username string, opts metav1.ListOptions) (*model.SecretList, error) {
	ret := &model.SecretList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	if username != "" {
		s.db = s.db.Where("username = ?", username)
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	name, _ := selector.RequiresExactMatch("name")

	d := s.db.Where(" name like ?", "%"+name+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}

func newSecrets(ds *dataStore) *secrets {
	return &secrets{ds.db}
}
