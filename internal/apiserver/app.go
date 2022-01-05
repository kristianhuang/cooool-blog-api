/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-go/internal/apiserver/config"
	"blog-go/internal/apiserver/options"
	"blog-go/pkg/app"
)

const commandDesc = `Welcome to use Blog-API`

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
		// TODO 需要做个 Log 包

		conf := config.NewConfig(opts)

		return Run(conf)
	}
}

func Run(conf *config.Config) error {
	server, err := createAPIServer(conf)
	if err != nil {
		return err
	}

	return server.BeforeRun().Run()
}
