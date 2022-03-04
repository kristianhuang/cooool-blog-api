/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package model

import (
	"time"

	"blog-api/pkg/auth"
	metav1 "blog-api/pkg/meta/v1"
	"blog-api/pkg/util/idutil"
	"gorm.io/gorm"
)

type AdminUser struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NickName          string `json:"nickname" gorm:"not null;unique;column:nickname;size:50;index:user;comment:昵称"`
	Password          string `json:"password" gorm:"not null;column:password;size:100;index:user;comment:密码"`
	Mobile            string `json:"mobile" gorm:"type:char(11);column:mobile;not null;index;comment:手机号"`
	Email             string `json:"email" gorm:"column:email;not null;size:70;comment:邮箱"`
	Status            uint8  `json:"status" gorm:"not null;default:2;size:1;comment:状态 1启用2禁用"`
	LoginedAt         int64  `json:"logined_at" gorm:"type:int(11);column:logined_at;not null;default:0;comment:最后登陆时间"`

	TotalPolicy int64 `json:"total_policy" gorm:"-"`

	metav1.DeleteMeta `json:",omitempty"`
}

type AdminUserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*AdminUser `json:"items"`
}

func (AdminUser) TableName() string {
	return "admin_user"
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `AdminUser` struct).
func (u *AdminUser) Compare(pwd string) error {
	return auth.Compare(u.Password, pwd)
}

func (u *AdminUser) AfterFind(tx *gorm.DB) error {
	if u.CreatedAt != 0 && u.UpdatedAt != 0 {
		u.CreatedAtFormat = time.Unix(u.CreatedAt, 0).Format("2006-01-02 15:04:05")
		u.UpdatedAtFormat = time.Unix(u.UpdatedAt, 0).Format("2006-01-02 15:04:05")
	}

	return nil
}

func (u *AdminUser) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = idutil.GetInstanceID(u.ID, "adminUser-")

	return tx.Save(u).Error
}
