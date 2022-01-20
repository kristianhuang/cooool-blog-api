/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"sync"

	"blog-api/internal/apiserver/store"
	"blog-api/internal/apiserver/store/mysql"
)

const (
	V1 = "/v1"
	// If you need more version...
)

var (
	once     sync.Once
	storeIns store.Factory
)

func init() {
	once.Do(func() {
		storeIns, _ = mysql.GetMysqlFactory(nil)
	})
}
