/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"encoding/json"

	genericoptions "blog-go/internal/pkg/options"
	"blog-go/pkg/cli/flag"
)

type Options struct {
	ServerRunOptions       *genericoptions.ServerRunOptions
	InsecureServingOptions *genericoptions.InsecureServingOptions
	FeatureOptions         *genericoptions.FeatureOptions
	MySQLOptions           *genericoptions.MySQLOptions
	RedisOptions           *genericoptions.RedisOptions
}

func NewOptions() *Options {
	return &Options{
		ServerRunOptions:       genericoptions.NewServerRunOptions(),
		InsecureServingOptions: genericoptions.NewInsecureServingOptions(),
		FeatureOptions:         genericoptions.NewFeatureOptions(),
		MySQLOptions:           genericoptions.NewMySQLOptions(),
		RedisOptions:           genericoptions.NewRedisOptions(),
	}
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.ServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServingOptions.AddFlags(fss.FlagSet("insecure serving"))
	o.FeatureOptions.AddFlags(fss.FlagSet("features"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	return fss
}

// Complete 设置需要默认值的选项
func (o Options) Complete() error {
	return nil
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
