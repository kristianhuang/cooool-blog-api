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
	metav1 "blog-api/pkg/meta/v1"
	"blog-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type delCollectionForm struct {
	Accounts []string `json:"accounts" validate:"required" form:"accounts" uri:"accounts"`
}

func (a *AdminUserController) DeleteCollection(c *gin.Context) {
	delCollectionForm := &delCollectionForm{}

	if err := c.ShouldBind(delCollectionForm); err != nil {
		response.Write(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	if err := delCollectionForm.validation(); err != nil {
		response.Write(c, err, nil)
		return
	}

	if err := a.srv.AdminUser().DeleteCollection(c, delCollectionForm.Accounts, metav1.DeleteOptions{Unscoped: true}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, nil)
}

func (f *delCollectionForm) validation() error {
	if err := validator.Struct(f); err != nil {
		return errors.WithCode(code.ErrValidation, err.(*validator.ValidationErrors).TranslateErrs()[0].Error())
	}

	return nil
}
