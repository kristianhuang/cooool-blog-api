/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"cooool-blog-api/internal/pkg/server"

	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	Mode        string   `json:"mode" mapstructure:"mode"`
	Health      bool     `json:"health" mapstructure:"health"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()

	return &ServerRunOptions{
		Mode:        defaults.Mode,
		Health:      defaults.Health,
		Middlewares: defaults.Middlewares,
	}
}

func (o *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Mode = o.Mode
	c.Health = o.Health
	c.Middlewares = o.Middlewares

	return nil
}

func (o *ServerRunOptions) Validate() []error {
	return []error{}
}

func (o *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&o.Health, "server.health", o.Health, "Add self readiness check and install /health router.")

	fs.StringVar(&o.Mode, "server.mode", o.Mode, "Server mode, Supported server mode: debug, test, release.")

	fs.StringSliceVar(&o.Middlewares, "server.middlewares", o.Middlewares, "List of allowed middlewares for server.")
}
