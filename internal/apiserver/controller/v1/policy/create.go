/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package policy

import (
	"cooool-blog-api/internal/pkg/middleware"
	"cooool-blog-api/internal/pkg/model"
	"cooool-blog-api/internal/pkg/response"
	metav1 "cooool-blog-api/pkg/meta/v1"

	"github.com/gin-gonic/gin"
)

type createForm struct {
	Name   string            `json:"name" form:"name" validate:"required,gt=6"`
	Policy model.AuthzPolicy `json:"policy,omitempty" form:"policy" validate:"omitempty"`
	Extend metav1.Extend     `json:",inline,omitempty" form:"extend" validate:"omitempty" label:"拓展字段"`
}

func (p *PolicyController) Create(c *gin.Context) {
	f := &createForm{}
	if err := p.Bind(c, f); err != nil {
		response.Write(c, err, nil)
		return
	}
	pm := f.applyTo()
	pm.UserName = c.GetString(middleware.UsernameKey)
	if err := p.srv.Policy().Create(c, pm, metav1.CreateOptions{}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, pm)
}

func (f *createForm) applyTo() *model.Policy {
	return &model.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Name:   f.Name,
			Extend: f.Extend,
		},
		Policy: f.Policy,
	}
}
