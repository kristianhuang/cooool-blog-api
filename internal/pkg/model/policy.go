/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package model

import (
	"blog-api/pkg/json"
	metav1 "blog-api/pkg/meta/v1"
	"blog-api/pkg/util/idutil"
	"github.com/ory/ladon"
	"gorm.io/gorm"
)

// AuthzPolicy defines iam policy type.
type AuthzPolicy struct {
	ladon.DefaultPolicy
}

type Policy struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserName string `json:"username" gorm:"not null;column:username;comment:账号;size:64;"`

	Policy AuthzPolicy `json:"policy,omitempty" gorm:"-"`

	PolicyShadow string `json:"-" gorm:"column:policy_shadow;"`
}

type PolicyList struct {
	metav1.ListMeta `json:",inline"`

	Items []*Policy `json:"items"`
}

// TableName maps to mysql table name.
func (p *Policy) TableName() string {
	return "policy"
}

func (p *AuthzPolicy) String() string {
	data, _ := json.Marshal(p)

	return string(data)
}

// BeforeCreate run before create database record.
func (p *Policy) BeforeCreate(tx *gorm.DB) error {
	if err := p.ObjectMeta.BeforeCreate(tx); err != nil {
		return err
	}

	p.PolicyShadow = p.Policy.String()

	return nil
}

// AfterCreate run after create database record.
func (p *Policy) AfterCreate(tx *gorm.DB) error {
	p.InstanceID = idutil.GetInstanceID(p.ID, "policy-")

	return tx.Save(p).Error
}

// BeforeUpdate run before update database record.
func (p *Policy) BeforeUpdate(tx *gorm.DB) error {
	if err := p.ObjectMeta.BeforeUpdate(tx); err != nil {
		return err
	}

	p.PolicyShadow = p.Policy.String()

	return nil
}

func (p *Policy) AfterFind(tx *gorm.DB) error {
	if err := p.ObjectMeta.AfterFind(tx); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(p.PolicyShadow), &p.Policy); err != nil {
		return err
	}

	return nil
}
