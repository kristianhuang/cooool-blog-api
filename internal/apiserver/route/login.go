/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import "github.com/gin-gonic/gin"

func Login(e *gin.Engine) {
	e.POST("/login", authStrategy.jwt.LoginHandler)
	e.POST("/logout", authStrategy.jwt.LogoutHandler)
	e.POST("/refresh", authStrategy.jwt.RefreshHandler)
}
