/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-api/internal/pkg/middleware"
	"blog-api/internal/pkg/middleware/auth"
)

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "blog-admin"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "blog-apiserver"
)

type loginInfo struct {
	Account  string `form:"account" json:"account" binding:"required,account"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func newJWTAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(username, password string) bool {
		// user, err := store.Client().AdminUser().Get(context.TODO(), username, metav1.GetOptions{})
		// if err != nil {
		// 	return false
		// }
		return true
	})
}
