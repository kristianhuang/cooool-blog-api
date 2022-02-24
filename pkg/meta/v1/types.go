/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Extend map[string]interface{}

func (e Extend) String() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// Merge merge extend fields from extendShadow.
func (e Extend) Merge(extendShadow string) Extend {
	var extend Extend

	_ = json.Unmarshal([]byte(extendShadow), &extend)
	for k, v := range extend {
		if _, ok := e[k]; !ok {
			e[k] = v
		}
	}
	return e
}

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"api_version,omitempty"`
}

type ListMeta struct {
	Total int64 `json:"total,omitempty"`
}

type ListOptions struct {
	TypeMeta       `json:",inline"`
	LabelSelector  string `json:"label_selector,omitempty" form:"label_selector"`
	FieldSelector  string `json:"field_selector,omitempty" form:"field_selector"`
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty"`
	Offset         *int64 `json:"offset,omitempty" form:"offset"`
	Limit          *int64 `json:"limit,omitempty" form:"limit"`
}

type ObjectMeta struct {
	ID        uint `json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
	CreatedAt int  `json:"created_at,omitempty" gorm:"type:int(15);not null;column:created_at;comment:创建时间;"`
	UpdatedAt int  `json:"updated_at,omitempty" gorm:"type:int(15);not null;column:updated_at;comment:更新时间;"`

	// 脱离于 db 的额外的拓展
	Extend       Extend `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ExtendShadow string `json:"-" gorm:"column:extend_shadow" validate:"omitempty"`
}

// BeforeCreate run before create database record.
func (obj *ObjectMeta) BeforeCreate(tx *gorm.DB) error {
	obj.ExtendShadow = obj.Extend.String()

	return nil
}

// BeforeUpdate run before update database record.
func (obj *ObjectMeta) BeforeUpdate(tx *gorm.DB) error {
	obj.ExtendShadow = obj.Extend.String()

	return nil
}

// AfterFind run after find to unmarshal a extend shadown string into metav1.Extend struct.
func (obj *ObjectMeta) AfterFind(tx *gorm.DB) error {
	if err := json.Unmarshal([]byte(obj.ExtendShadow), &obj.Extend); err != nil {
		return err
	}

	return nil
}

type DeleteOptions struct {
	TypeMeta `json:",inline"`
	Unscoped bool `json:"unscoped"`
}

type CreateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dry_run,omitempty"`
}

type UpdateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dry_run,omitempty"`
}

type PatchOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dry_run,omitempty"`
	Force    bool     `json:"force,omitempty"`
}

type AuthorizeOptions struct {
	TypeMeta `json:",inline"`
}

// GetOptions is the standard query options to the standard REST get call.
type GetOptions struct {
	TypeMeta `json:",inline"`
}

type TableOptions struct {
	TypeMeta  `json:",inline"`
	NoHeaders bool `json:"-"`
}
