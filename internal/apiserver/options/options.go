/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	genericoptions "blog-api/internal/pkg/options"
	"blog-api/internal/pkg/server"
	"blog-api/pkg/flag"
	"blog-api/pkg/json"
	"blog-api/pkg/rollinglog"
	"blog-api/pkg/util/idutil"
	"blog-api/pkg/validator"
)

// TODO https options
type Options struct {
	ServerRunOptions       *genericoptions.ServerRunOptions       `json:"server" mapstructure:"server"`
	InsecureServingOptions *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	FeatureOptions         *genericoptions.FeatureOptions         `json:"feature" mapstructure:"feature"`
	MySQLOptions           *genericoptions.MySQLOptions           `json:"mysql" mapstructure:"mysql"`
	RedisOptions           *genericoptions.RedisOptions           `json:"redis" mapstructure:"redis"`
	JwtOptions             *genericoptions.JwtOptions
	Log                    *rollinglog.Options `json:"log" mapstructure:"log"`
	Validator              *validator.Options  `json:"validator" mapstructure:"validator"`
}

func NewOptions() *Options {
	return &Options{
		ServerRunOptions:       genericoptions.NewServerRunOptions(),
		InsecureServingOptions: genericoptions.NewInsecureServingOptions(),
		FeatureOptions:         genericoptions.NewFeatureOptions(),
		MySQLOptions:           genericoptions.NewMySQLOptions(),
		RedisOptions:           genericoptions.NewRedisOptions(),
		JwtOptions:             genericoptions.NewJwtOptions(),
		Log:                    rollinglog.NewOptions(),
		Validator:              validator.NewOptions(),
	}
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.ServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServingOptions.AddFlags(fss.FlagSet("insecure serving"))
	o.FeatureOptions.AddFlags(fss.FlagSet("features"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.Log.AddFlags(fss.FlagSet("log"))
	o.Validator.AddFlags(fss.FlagSet("validator"))
	o.JwtOptions.AddFlags(fss.FlagSet("jwt"))

	return fss
}

// Complete 设置需要默认值的选项
func (o *Options) Complete() error {
	if o.JwtOptions.Key == "" {
		o.JwtOptions.Key = idutil.NewSecretKey()
	}

	return nil
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
