/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import "encoding/json"

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
	ID uint64 `json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
	// 脱离于 db 的额外的拓展
	Extend       Extend `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ExtendShadow string `json:"-" gorm:"column:extendShadow" validate:"omitempty"`
	CreatedAt    int    `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt    int    `json:"updated_at,omitempty" gorm:"column:updated_at"`
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

type TableOptions struct {
	TypeMeta  `json:",inline"`
	NoHeaders bool `json:"-"`
}
