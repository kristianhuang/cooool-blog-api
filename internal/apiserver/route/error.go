/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	errorcontroller "blog-api/internal/apiserver/controller/v1/error"
	"github.com/gin-gonic/gin"
)

func Error(e *gin.Engine) {
	// Page not found route.
	errController := errorcontroller.NewErrorController()
	e.NoRoute(errController.PageNotFound)
}
