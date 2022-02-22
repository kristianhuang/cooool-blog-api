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
	"blog-api/pkg/errors"
	metav1 "blog-api/pkg/meta/v1"
	"blog-api/pkg/validator"
	"github.com/gin-gonic/gin"

	validation_util "blog-api/pkg/validator/util"
)

type CreateForm struct {
	Account  string `json:"account" form:"account" validate:"required,gt=6" label:"账号"`
	NickName string `json:"nick_name" form:"nick_name" validate:"required" label:"昵称"`
	Password string `json:"password" form:"password" validate:"required,gt=0" label:"密码"`
	Mobile   string `json:"mobile" form:"mobile" validate:"required,isMobile" label:"手机号"`
	Email    string `json:"email" form:"email" validate:"required,email" label:"邮箱"`
	Status   uint8  `json:"status" form:"status" validate:"required,oneof=1 2" label:"状态"`
}

func (a *AdminUserController) Create(c *gin.Context) {
	formData := &CreateForm{}
	if err := c.ShouldBind(formData); err != nil {
		response.Write(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	v := validator.New("zh", "label")
	if err := v.RegisterValidation("isMobile", "请输入有效手机号码", validation_util.Mobile); err != nil {
		response.Write(c, errors.WithCode(code.ErrRegisterValidation, err.Error()), nil)

		return
	}

	if err := v.Struct(formData); err != nil {
		response.Write(c, errors.WithCode(code.ErrValidation, err.(*validator.ValidationErrors).TranslateErrs()[0].Error()), nil)

		return
	}

	au := &model.AdminUser{
		Account:  formData.Account,
		NickName: formData.NickName,
		Password: formData.Password,
		Mobile:   formData.Mobile,
		Email:    formData.Email,
		Status:   formData.Status,
	}

	if err := a.srv.AdminUser().Create(c, au, metav1.CreateOptions{}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, au)
}
