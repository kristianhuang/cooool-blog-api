/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package model

import (
	metav1 "blog-api/pkg/meta/v1"
	"blog-api/pkg/util/idutil"
	"gorm.io/gorm"
)

type Secret struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Account           string `json:"account" gorm:"column:account;size:50;comment:账号信息"`
	SecretID          string `json:"secret_id" gorm:"column:secret_id;comment:secret_id;"`
	SecretKey         string `json:"secret_key" gorm:"column:secret_key;comment:secret_key;"`

	Expires     int64  `json:"expires" gorm:"column:expires;comment:有效期;"`
	Description string `json:"description" gorm:"column:description;comment:描述;"`
}

type SecretList struct {
	metav1.ListMeta `json:",inline"`

	Items []*Secret `json:"items"`
}

func (s *Secret) TableName() string {
	return "secret"
}

// AfterCreate run after create database record.
func (s *Secret) AfterCreate(tx *gorm.DB) error {
	s.InstanceID = idutil.GetInstanceID(s.ID, "secret-")

	return tx.Save(s).Error
}
