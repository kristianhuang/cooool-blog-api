/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package response

import (
	"net/http"

	"cooool-blog-api/internal/pkg/code"
	"cooool-blog-api/pkg/errors"
	"cooool-blog-api/pkg/rollinglog"

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

// SuccessResp defines the return messages when a success occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type SuccessResp struct {
	// Code default the code.ErrSuccess.
	Code int `json:"code"`
	// Data is responses data.
	Data interface{} `json:"data,omitempty"`
}

// Write an error, or the response data into http response body.
// It uses errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func Write(c *gin.Context, err error, data interface{}) {
	// TODO
	//  - 自定义错误信息
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

	c.JSON(http.StatusOK, SuccessResp{
		Code: code.ErrSuccess,
		Data: data,
	})
}
