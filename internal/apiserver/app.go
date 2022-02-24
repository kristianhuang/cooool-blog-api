/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-api/internal/apiserver/config"
	"blog-api/internal/apiserver/options"
	"blog-api/pkg/app"
	log "blog-api/pkg/rollinglog"
)

const commandDesc = `Welcome to use Blog-API-SERVER`

func NewApp(use string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		use,
		"api-server",
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithLong(commandDesc),
		app.WithRunFunc(createRunFunc(opts)),
	)

	return application
}

func createRunFunc(opts *options.Options) app.RunFunc {
	return func(use string) error {
		log.Init(opts.Log)
		defer log.Flush()

		return Run(config.NewConfig(opts))
	}
}

func Run(conf *config.Config) error {
	server, err := createAPIServer(conf)
	if err != nil {
		return err
	}

	return server.BeforeRun().Run()
}
