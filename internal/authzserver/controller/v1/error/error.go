/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package error

import (
	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/internal/pkg/response"
	"cooool-blog-api/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ErrorController struct {
}

func NewErrorController() *ErrorController {
	return &ErrorController{}
}

func (ec *ErrorController) PageNotFound(c *gin.Context) {
	response.Write(c, errors.WithCode(code.ErrPageNotFound, "Page not found."), nil)
}
