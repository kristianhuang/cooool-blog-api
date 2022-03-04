/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authorize

import (
	"blog-api/internal/authzserver/authorization"
	"blog-api/internal/authzserver/authorization/authorizer"
	"blog-api/internal/pkg/bind"
	"blog-api/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
)

type AuthzController struct {
	store authorizer.PolicyGetter

	*bind.BindData
}

func NewAuthzController(store authorizer.PolicyGetter) *AuthzController {
	return &AuthzController{
		store:    store,
		BindData: bind.NewBindData(),
	}
}

func (a *AuthzController) Authorize(c *gin.Context) {
	var r ladon.Request
	if err := a.Bind(c, &r); err != nil {
		response.Write(c, err, nil)
		return
	}

	auth := authorization.NewAuthorizer(authorizer.NewAuthorization(a.store))

	if r.Context == nil {
		r.Context = ladon.Context{}
	}

	r.Context["username"] = c.GetString("username")
	rsp := auth.Authorize(&r)

	response.Write(c, nil, rsp)
}
