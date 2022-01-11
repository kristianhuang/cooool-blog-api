/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package fields

import (
	"sort"
	"strings"
)

type Fields interface {
	// Has return whether the field exists.
	Has(field string) (ok bool)
	// Get return the value for the filed.
	Get(field string) (val string)
}

type Set map[string]string

func (s Set) String() string {
	selector := make([]string, 0, len(s))

	for key, val := range s {
		selector = append(selector, key+"="+val)
	}
	sort.StringSlice(selector).Sort()
	return strings.Join(selector, ",")
}

func (s Set) Has(field string) bool {
	_, ok := s[field]
	return ok
}

func (s Set) Get(field string) string {

	return s[field]
}

func (s Set) AsSelector() Selector {
	return SelectorFromSet(s)
}
