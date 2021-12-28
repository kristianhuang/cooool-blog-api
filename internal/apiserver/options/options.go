/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	genericoptions "blog-go/internal/pkg/options"
	"blog-go/pkg/cli/flag"
	"encoding/json"
)

type Options struct {
	APISServerOptions *genericoptions.APIServerOptions
	MySQLOptions      *genericoptions.MySQLOptions
	RedisOptions      *genericoptions.RedisOptions
}

func NewOptions() *Options {
	return &Options{
		APISServerOptions: genericoptions.NewServerOptions(),
		MySQLOptions:      genericoptions.NewMySQLOptions(),
		RedisOptions:      genericoptions.NewRedisOptions(),
	}
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.APISServerOptions.AddFlags(fss.FlagSet("api-server"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
