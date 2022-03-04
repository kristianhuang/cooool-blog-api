/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package bind

import (
	"blog-api/internal/pkg/code"
	"blog-api/pkg/errors"
	"blog-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type BindData struct{}

func NewBindData() *BindData {
	return &BindData{}
}

// BindJson bind json data and validate.
func (b *BindData) BindJson(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindJSON(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := b.validate(data); err != nil {
		return err
	}

	return nil
}

// BindQuery bind query data and validate.
func (b *BindData) BindQuery(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindQuery(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := b.validate(data); err != nil {
		return err
	}

	return nil
}

// BindUri bind uri data and validate.
func (b *BindData) BindUri(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindUri(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := b.validate(data); err != nil {
		return err
	}

	return nil
}

// Bind bind data and validate.
func (b *BindData) Bind(c *gin.Context, data interface{}) error {
	if err := c.ShouldBind(data); err != nil {
		return errors.WithCode(code.ErrBind, err.Error())
	}

	if err := b.validate(data); err != nil {
		return err
	}

	return nil
}

func (b *BindData) validate(data interface{}) error {
	if err := validator.Struct(data); err != nil {
		return errors.WithCode(code.ErrValidation, err.(*validator.ValidationErrors).TranslateErrs()[0].Error())
	}

	return nil
}
