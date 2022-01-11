/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	metav1 "blog-api/pkg/meta/v1"
)

type AdminUser struct {
	metav1.ObjectMeta `json:"meta_data,omitempty"`
	ID                int64  `json:"id" gormutil:"primaryKey;autoIncrement"`
	Account           string `json:"account" gormutil:"not null"`
}

type AdminUserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*AdminUser `json:"items"`
}

func (u *AdminUser) TableName() string {
	return "admin_user"
}
