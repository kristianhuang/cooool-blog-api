/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package config

import "cooool-blog-api/internal/apiserver/options"

type Config struct {
	*options.Options
}

func NewConfig(options *options.Options) *Config {
	return &Config{Options: options}
}
