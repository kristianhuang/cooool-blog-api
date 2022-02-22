/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/response"
	"blog-api/pkg/errors"
	"blog-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type DelForm struct {
	Account string `json:"account" validate:"required,gt=6" form:"account" uri:"account"`
}

func (a *AdminUserController) Delete(c *gin.Context) {
	data := &DelForm{}
	if err := c.ShouldBind(data); err != nil {
		response.Write(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if err := validator.Struct(data); err != nil {
		response.Write(c, errors.WithCode(code.ErrValidation, err.(*validator.ValidationErrors).TranslateErrs()[0].Error()), nil)

		return
	}

	response.Write(c, nil, nil)
}
