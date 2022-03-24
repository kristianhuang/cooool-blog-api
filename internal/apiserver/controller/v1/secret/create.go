/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package secret

import (
	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/internal/pkg/middleware"
	"cooool-blog-api/internal/pkg/model"
	"cooool-blog-api/internal/pkg/response"
	"cooool-blog-api/pkg/errors"
	metav1 "cooool-blog-api/pkg/meta/v1"
	"cooool-blog-api/pkg/util/idutil"

	"github.com/gin-gonic/gin"
)

type createForm struct {
	Name string `json:"name" form:"name" validate:"omitempty"`

	// Required: true
	Expires     int64  `json:"expires"   form:"expires"      validate:"omitempty"`
	Description string `json:"description" form:"description"  validate:"omitempty"`

	Extend metav1.Extend `json:",inline,omitempty" form:"extend" validate:"omitempty" label:"拓展字段"`
}

const maxSecretCount = 10

func (s *SecretController) Create(c *gin.Context) {
	f := &createForm{}
	if err := s.Bind(c, f); err != nil {
		response.Write(c, err, nil)
		return
	}

	username := c.GetString(middleware.UsernameKey)
	secrets, err := s.srv.Secrets().List(c, username, metav1.ListOptions{})

	if err != nil {
		response.Write(c, err, nil)
		return
	}

	if secrets.TotalCount >= maxSecretCount {
		response.Write(c, errors.WithCode(code.ErrReachMaxCount, "secret count: %d", secrets.TotalCount), nil)
		return
	}

	if err := s.srv.Secrets().Create(c, f.applyTo(username), metav1.CreateOptions{}); err != nil {
		response.Write(c, err, nil)
		return
	}

	response.Write(c, nil, f)
}

func (f *createForm) applyTo(username string) *model.Secret {
	return &model.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:   f.Name,
			Extend: f.Extend,
		},
		UserName:    username,
		SecretID:    idutil.NewSecretID(),
		SecretKey:   idutil.NewSecretKey(),
		Expires:     f.Expires,
		Description: f.Description,
	}
}
