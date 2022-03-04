/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/model"
	"blog-api/internal/pkg/response"
	"blog-api/pkg/auth"
	"blog-api/pkg/errors"
	metav1 "blog-api/pkg/meta/v1"
	"blog-api/pkg/validator"
	validationutil "blog-api/pkg/validator/util"
	"github.com/gin-gonic/gin"
)

type updateForm struct {
	NickName string        `json:"nickname,omitempty" form:"nickname" validate:"omitempty,gt=2" label:"昵称"`
	Password string        `json:"password,omitempty" form:"password" validate:"omitempty,gt=6" label:"密码"`
	Mobile   string        `json:"mobile,omitempty" form:"mobile" validate:"omitempty,isMobile" label:"手机号"`
	Email    string        `json:"email,omitempty" form:"email" validate:"omitempty,email" label:"邮箱"`
	Status   uint8         `json:"status,omitempty" form:"status" validate:"omitempty,oneof=1 2" label:"状态"`
	Extend   metav1.Extend `json:",inline,omitempty" form:"extend" validate:"omitempty" label:"拓展字段"`
}

func (a *AdminUserController) Update(c *gin.Context) {
	f := &updateForm{}
	if err := c.ShouldBind(f); err != nil {
		response.Write(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	if err := f.validate(); err != nil {
		response.Write(c, err, nil)
		return
	}

	au, err := a.srv.AdminUser().Get(c, c.Param("name"), metav1.GetOptions{})
	if err != nil {
		response.Write(c, err, nil)
		return
	}

	f.applyTo(au)
	if err := a.srv.AdminUser().Update(c, au, metav1.UpdateOptions{}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, au)
}

func (f *updateForm) applyTo(au *model.AdminUser) {
	// If password is not empty, and the password is not equal to the old password, set the new password.
	if len(f.Password) > 0 {
		if err := auth.Compare(au.Password, f.Password); err != nil {
			au.Password, _ = auth.Encrypt(f.Password)
		}
	}

	au.NickName = f.NickName
	au.Mobile = f.Mobile
	au.Email = f.Email
	au.Extend = f.Extend
	au.Status = f.Status
}

func (f *updateForm) validate() error {
	v := validator.New("zh", "label")
	if err := v.RegisterValidation("isMobile", "请输入有效手机号码", validationutil.Mobile); err != nil {
		return errors.WithCode(code.ErrRegisterValidation, err.Error())
	}

	if err := v.Struct(f); err != nil {
		return errors.WithCode(code.ErrValidation, err.(*validator.ValidationErrors).TranslateErrs()[0].Error())
	}

	return nil
}
