/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/internal/pkg/response"
	"cooool-blog-api/pkg/auth"
	"cooool-blog-api/pkg/errors"
	metav1 "cooool-blog-api/pkg/meta/v1"

	"github.com/gin-gonic/gin"
)

type changePasswordForm struct {
	OldPassword string `json:"old_password" form:"old_password" validate:"required,gt=6" label:"旧密码"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required,gt=6" label:"新密码"`
}

func (a *AdminUserController) ChangePassword(c *gin.Context) {
	f := &changePasswordForm{}
	if err := a.Bind(c, f); err != nil {
		response.Write(c, err, nil)
		return
	}

	au, err := a.srv.AdminUser().Get(c, c.Param("username"), metav1.GetOptions{})
	if err != nil {
		response.Write(c, err, nil)
		return
	}

	if err := au.Compare(f.OldPassword); err != nil {
		response.Write(c, errors.WithCode(code.ErrPasswordIncorrect, err.Error()), nil)
		return
	}

	au.Password, _ = auth.Encrypt(f.NewPassword)
	if err := a.srv.AdminUser().ChangePassword(c, au); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, nil)
}
