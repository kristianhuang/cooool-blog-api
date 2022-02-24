/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package model

import (
	metav1 "blog-api/pkg/meta/v1"
	"gorm.io/plugin/soft_delete"
)

type AdminUser struct {
	metav1.ObjectMeta `json:"meta_data,omitempty"`
	Account           string                `json:"account" gorm:"not null;column:account;size:50;comment:账号信息"`
	NickName          string                `json:"nick_name" gorm:"not null;column:nick_name;size:50;index:user;comment:昵称"`
	Password          string                `json:"password" gorm:"not null;column:password;size:50;index:user;comment:密码"`
	Mobile            string                `json:"mobile" gorm:"type:char(11);column:mobile;not null;index;comment:手机号"`
	Email             string                `json:"email" gorm:"column:email;not null;size:70;comment:邮箱"`
	Status            uint8                 `json:"status" gorm:"not null;default:2;size:1;comment:状态 1启用2禁用"`
	LoginedAt         int                   `json:"logined_at" gorm:"type:int(11);column:logined_at;not null;default:0;comment:最后登陆时间"`
	DeletedAt         soft_delete.DeletedAt `json:"deleted_at,omitempty" gorm:"type:int(15);column:deleted_at;comment:删除时间"`
}

type AdminUserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*AdminUser `json:"items"`
}

func (AdminUser) TableName() string {
	return "admin_user"
}
