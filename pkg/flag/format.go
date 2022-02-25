/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package flag

import (
	"log"
	"strings"

	"github.com/spf13/pflag"
)

// WordReplaceNormalizeFunc 转化非标准格式 flag 为标准格式
func WordReplaceNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators.
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		name := strings.ReplaceAll(name, "_", "-")
		log.Printf("%s is DEPRECATED and will be removed in a future version. Use %s instead.", name, name)
		return pflag.NormalizedName(name)
	}
	return pflag.NormalizedName(name)
}
