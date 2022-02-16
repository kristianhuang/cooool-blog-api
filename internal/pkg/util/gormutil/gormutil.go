/*
 * Copyright 2021 KristianHuang <KristianHuangyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package gormutil

const (
	DefaultLimit = 25
)

type LimitAndOffset struct {
	Offset int
	Limit  int
}

// Unpointer fill LimitAndOffset with default val, if offset/limit is nil.
func Unpointer(offset *int64, limit *int64) *LimitAndOffset {
	var o, l int = 0, DefaultLimit

	if offset != nil {
		o = int(*offset)
	}

	if limit != nil {
		l = int(*limit)
	}

	return &LimitAndOffset{
		Offset: o,
		Limit:  l,
	}
}
