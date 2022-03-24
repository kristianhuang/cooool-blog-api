/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
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

type adminUser struct {
	db *gorm.DB
}

func newAdminUser(ds *dataStore) *adminUser {
	return &adminUser{ds.db}
}

func (u *adminUser) Create(ctx context.Context, adminUserModel *model.AdminUser, opts metav1.CreateOptions) error {
	return u.db.Create(adminUserModel).Error
}

func (u *adminUser) Update(ctx context.Context, adminUserModel *model.AdminUser, opts metav1.UpdateOptions) error {
	return u.db.Save(adminUserModel).Error
}

func (u *adminUser) List(cxt context.Context, opts metav1.ListOptions) (*model.AdminUserList, error) {
	userList := &model.AdminUserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)
	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, _ := selector.RequiresExactMatch("name")
	d := u.db.Where("name like ?", "%"+username+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&userList.Items).
		Offset(-1).
		Limit(-1).
		Count(&userList.TotalCount)

	return userList, d.Error
}

func (u *adminUser) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	pol := newPolicy(&dataStore{u.db})
	if err := pol.DeleteByAdminUser(ctx, username, opts); err != nil {
		return err
	}

	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	err := u.db.Where("name = ?", username).Delete(&model.AdminUser{}).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (u *adminUser) Get(ctx context.Context, username string, opts metav1.GetOptions) (*model.AdminUser, error) {
	au := &model.AdminUser{}
	err := u.db.Where("name = ?", username).First(&au).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return au, nil
}

func (u *adminUser) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	// pol := newPolicy(&dataStore{u.db})
	// if err := pol.DeleteCollectionByAdminUser(ctx, usernames, opts); err != nil {
	// 	return err
	// }

	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	return u.db.Where("username in (?)", usernames).Delete(&model.AdminUser{}).Error
}
