/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package watcher

import (
	genericapiserver "cooool-blog-api/internal/pkg/server"
	"cooool-blog-api/internal/watcher/config"
	"cooool-blog-api/internal/watcher/options"
	"cooool-blog-api/pkg/app"
	log "cooool-blog-api/pkg/rollinglog"
)

const commandDesc = `cooool-blog Watcher is a pluggable watcher service used to do some periodic work like cron job. 
But the difference with cron job is blog-watcher also support sleep some duration after previous job done.
`

// NewApp creates an App object with default parameters.
func NewApp(use string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		use,
		"watcher-server",
		app.WithOptions(opts),
		app.WithLong(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
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

// Run runs the specified pump server. This should never exit.
func Run(cfg *config.Config) error {
	go genericapiserver.ServeHealthCheck(cfg.HealthCheckPath, cfg.HealthCheckAddress)

	return createWatcherServer(cfg).BeforeRun().Run()
}
