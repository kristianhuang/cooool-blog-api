/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	errorcontroller "cooool-blog-api/internal/authzserver/controller/v1/error"

	"github.com/gin-gonic/gin"
)

func Error(e *gin.Engine) {
	// Page not found route.
	errController := errorcontroller.NewErrorController()
	e.NoRoute(errController.PageNotFound)
}
