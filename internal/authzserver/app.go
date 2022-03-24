/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authzserver

import (
	"cooool-blog-api/internal/authzserver/config"
	"cooool-blog-api/internal/authzserver/options"
	"cooool-blog-api/pkg/app"
	log "cooool-blog-api/pkg/rollinglog"
)

const commandDesc = `Authorization server to run ladon policies which can protecting your resources.
It is written inspired by AWS IAM policiis.

Find more ladon information at:
    https://github.com/ory/ladon`

func NewApp(use string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		use,
		"auth-server",
		app.WithOptions(opts),
		app.WithLong(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(createRunFunc(opts)),
	)

	return application
}

func createRunFunc(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}

// Run runs the specified AuthzServer. This should never exit.
func Run(cfg *config.Config) error {
	server, err := createAuthzServer(cfg)
	if err != nil {
		return err
	}

	return server.BeforeRun().Run()
}
