/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"blog-api/internal/authzserver/analytics"
	genericoptions "blog-api/internal/pkg/options"
	"blog-api/internal/pkg/server"
	cliflag "blog-api/pkg/flag"
	"blog-api/pkg/json"
	"blog-api/pkg/rollinglog"
)

type Options struct {
	RPCServer string `json:"rpcserver" mapstructure:"rpcserver"`
	ClientCA  string `json:"client-ca-file" mapstructure:"client-ca-file"`

	ServerRunOptions *genericoptions.ServerRunOptions       `json:"server" mapstructure:"server"`
	InsecureServing  *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	// https
	SecureServing    *genericoptions.SecureServingOptions `json:"secure" mapstructure:"secure"`
	RedisOptions     *genericoptions.RedisOptions         `json:"redis" mapstructure:"redis"`
	FeatureOptions   *genericoptions.FeatureOptions       `json:"feature" mapstructure:"feature"`
	Log              *rollinglog.Options                  `json:"log" mapstructure:"log"`
	AnalyticsOptions *analytics.AnalyticsOptions          `json:"analytics"      mapstructure:"analytics"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		RPCServer:        "127.0.0.1:8081",
		ClientCA:         "",
		ServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:  genericoptions.NewInsecureServingOptions(),
		SecureServing:    genericoptions.NewSecureServingOptions(),
		RedisOptions:     genericoptions.NewRedisOptions(),
		FeatureOptions:   genericoptions.NewFeatureOptions(),
		Log:              rollinglog.NewOptions(),
		AnalyticsOptions: analytics.NewAnalyticsOptions(),
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.ServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.AnalyticsOptions.AddFlags(fss.FlagSet("analytics"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.FeatureOptions.AddFlags(fss.FlagSet("features"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.Log.AddFlags(fss.FlagSet("logs"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	fs.StringVar(&o.RPCServer, "rpcserver", o.RPCServer, "The address of iam rpc server. "+
		"The rpc server can provide all the secrets and policies to use.")
	fs.StringVar(&o.ClientCA, "client-ca-file", o.ClientCA, ""+
		"If set, any request presenting a client certificate signed by one of "+
		"the authorities in the client-ca-file is authenticated with an identity "+
		"corresponding to the CommonName of the client certificate.")

	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Complete set default Options.
func (o *Options) Complete() error {
	return o.SecureServing.Complete()
}
