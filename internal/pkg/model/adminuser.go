/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package model

import (
	metav1 "blog-api/pkg/meta/v1"
)

type AdminUser struct {
	metav1.ObjectMeta `json:"meta_data,omitempty"`
	Account           string `json:"account" gorm:"not null"`
	NickName          string `json:"nick_name" gorm:"not null"`
	Password          string `json:"password" gorm:""`
}

type AdminUserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*AdminUser `json:"items"`
}

func (u *AdminUser) TableName() string {
	return "admin_user"
}
