/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package stringutil

import "unicode/utf8"

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func In(str string, array []string) bool {
	return FindIndex(array, str) > -1
}

func FindIndex(array []string, str string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}
