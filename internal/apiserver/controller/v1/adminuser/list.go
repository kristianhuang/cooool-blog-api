/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"cooool-blog-api/internal/pkg/response"
	metav1 "cooool-blog-api/pkg/meta/v1"

	"github.com/gin-gonic/gin"
)

type listForm struct {
	LabelSelector string `json:"label_selector,omitempty" uri:"label_selector"`
	FieldSelector string `json:"field_selector,omitempty" uri:"field_selector"`
	Offset        *int64 `json:"offset,omitempty" uri:"offset" validate:"omitempty,min=1,numeric"`
	Limit         *int64 `json:"limit,omitempty" uri:"limit" validate:"omitempty,min=1,numeric"`
	Status        uint8  `json:"status" uri:"status" validate:"omitempty,oneof=1 2"`
}

func (a *AdminUserController) List(c *gin.Context) {
	f := &listForm{}
	if err := a.BindUri(c, f); err != nil {
		response.Write(c, err, nil)
		return
	}

	aus, err := a.srv.AdminUser().List(c, f.applyTo())
	if err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, aus)
}

func (f *listForm) applyTo() metav1.ListOptions {
	return metav1.ListOptions{
		LabelSelector: f.LabelSelector,
		FieldSelector: f.FieldSelector,
		Offset:        f.Offset,
		Limit:         f.Limit,
	}
}
