/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package response

import (
	"net/http"

	"blog-api/pkg/errors"
	"blog-api/pkg/rollinglog"
	"github.com/gin-gonic/gin"
)

// ErrResp defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type ErrResp struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external.
	Message string `json:"message"`

	// Reference returns the reference document, which maybe useful to solve this error.
	Reference string `json:"reference"`
}

// Write an error, or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func Write(c *gin.Context, err error, data interface{}) {
	if err != nil {
		rollinglog.L(c).Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResp{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
