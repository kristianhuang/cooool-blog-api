/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package util

import (
	"regexp"

	go_validator "github.com/go-playground/validator/v10"
)

// Mobile 验证字段是否为有效电话
func Mobile(fl go_validator.FieldLevel) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)

	return reg.MatchString(fl.Field().String())
}
