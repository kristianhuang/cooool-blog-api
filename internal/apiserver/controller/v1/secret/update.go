/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package secret

import (
	"blog-api/internal/pkg/code"
	"blog-api/internal/pkg/middleware"
	"blog-api/internal/pkg/model"
	"blog-api/internal/pkg/response"
	"blog-api/pkg/errors"
	metav1 "blog-api/pkg/meta/v1"
	"blog-api/pkg/util/idutil"
	"github.com/gin-gonic/gin"
)

type updateForm struct {
	// Required: true
	Expires     int64  `json:"expires"   form:"expires"      validate:"omitempty"`
	Description string `json:"description" form:"description"  validate:"omitempty"`

	Extend metav1.Extend `json:",inline,omitempty" form:"extend" validate:"omitempty" label:"拓展字段"`
}

func (s *SecretController) Update(c *gin.Context) {
	f := &updateForm{}
	if err := s.Bind(c, f); err != nil {
		response.Write(c, err, nil)
		return
	}

	username := c.GetString(middleware.UsernameKey)
	name := c.Param("name")
	secret, err := s.srv.Secrets().Get(c, username, name, metav1.GetOptions{})
	if err != nil {
		response.Write(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	secret.Expires = f.Expires
	secret.Description = f.Description
	secret.Extend = f.Extend

	if err := s.srv.Secrets().Update(c, secret, metav1.UpdateOptions{}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, secret)
}

func (f *updateForm) applyTo(username string) *model.Secret {
	return &model.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Extend: f.Extend,
		},
		UserName:    username,
		SecretID:    idutil.NewSecretID(),
		SecretKey:   idutil.NewSecretKey(),
		Expires:     f.Expires,
		Description: f.Description,
	}
}
