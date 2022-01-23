/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package route

import (
	"blog-api/internal/apiserver/store"
)

const (
	V1 = "/v1"
	// If you need more version...
)

var (
	storeIns store.Factory
)

func SetStoreIns(s store.Factory) {
	storeIns = s
}
