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

type updateForm struct {
	Username string            `json:"username" form:"username" validate:"omitempty"`
	Policy   model.AuthzPolicy `json:"policy,omitempty" form:"policy" validate:"omitempty"`
	Extend   metav1.Extend     `json:",inline,omitempty" form:"extend" validate:"omitempty" label:"拓展字段"`
}

func (p *PolicyController) Update(c *gin.Context) {
	f := &updateForm{}
	if err := p.Bind(c, f); err != nil {
		response.Write(c, err, nil)
		return
	}

	pol, err := p.srv.Policy().Get(c, c.GetString(middleware.UsernameKey), c.Param("name"), metav1.GetOptions{})
	if err != nil {
		response.Write(c, err, nil)
		return
	}

	pol.Policy = f.Policy
	pol.Extend = f.Extend

	if err := p.srv.Policy().Update(c, pol, metav1.UpdateOptions{}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, pol)
}
