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
		"apiserver",
		app.WithFlags(opts),
		app.WithLong(commandDesc),
		app.WithRunFunc(createRunFunc(opts)),
	)

	return application
}

func createRunFunc(opts *options.Options) app.RunFunc {
	// TODO 需要做个 Log 包
	return func(use string) error {
		conf := config.NewConfig(opts)

		return run(conf)
	}
}

func run(cfg *config.Config) error {
	return nil
	// server, err := createAPIServer(cfg)
	// if err != nil {
	// 	return err
	// }
	//
	// return server.PrepareRun().Run()
}
