/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package adminuser

import (
	"net/http"

	"blog-api/internal/pkg/model"
	metav1 "blog-api/pkg/meta/v1"
	log "blog-api/pkg/rollinglog"
	"blog-api/pkg/validator"
	"github.com/gin-gonic/gin"

	validation_util "blog-api/pkg/validator/util"
)

type createForm struct {
	Account  string `json:"account" form:"account" validate:"required,gt=6" label:"账号"`
	NickName string `json:"nick_name" form:"nick_name" validate:"required" label:"昵称"`
	Password string `json:"password" form:"password" validate:"required,gt=0" label:"密码"`
	Mobile   string `json:"mobile" form:"mobile" validate:"required,isMobile" label:"手机号"`
	Email    string `json:"email" form:"email" validate:"required,email" label:"邮箱"`
	Status   uint8  `json:"status" form:"status" validate:"required,oneof=1 2" label:"状态"`
}

func (a *AdminUserController) Create(c *gin.Context) {
	formData := &createForm{}
	if err := c.ShouldBind(formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	v := validator.New("zh", "label")
	if err := v.RegisterValidation("isMobile", "请输入有效手机号码", validation_util.Mobile); err != nil {
		log.L(c).Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := v.Struct(formData); err != nil {
		log.L(c).Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
