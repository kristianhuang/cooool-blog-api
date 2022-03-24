/*
 * Copyright 2021 KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"cooool-blog-api/internal/pkg/server"

	"github.com/spf13/pflag"
)

type FeatureOptions struct {
	EnableProfiling bool `json:"enable_profiling" mapstructure:"enable_profiling"`
	EnableMetrics   bool `json:"enable_metrics" mapstructure:"enable_metrics"`
}

func NewFeatureOptions() *FeatureOptions {
	defaults := server.NewConfig()

	return &FeatureOptions{
		EnableProfiling: defaults.EnableProfiling,
		EnableMetrics:   defaults.EnableMetrics,
	}
}

func (o *FeatureOptions) ApplyTo(c *server.Config) error {
	c.EnableProfiling = o.EnableProfiling
	c.EnableMetrics = o.EnableMetrics
	return nil
}

func (o *FeatureOptions) Validate() []error {
	return []error{}
}

func (o *FeatureOptions) AddFlags(fs *pflag.FlagSet) {

	fs.BoolVar(&o.EnableProfiling, "feature.profiling", o.EnableProfiling, "Enable profiling via web interface host:port/debug/pprof")

	fs.BoolVar(&o.EnableMetrics, "feature.metrics", o.EnableMetrics, "Enables metrics on the apiserver at /metrics")
}
