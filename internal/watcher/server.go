/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package watcher

import (
	"time"

	"cooool-blog-api/internal/apiserver/store/mysql"
	genericoptions "cooool-blog-api/internal/pkg/options"
	"cooool-blog-api/internal/watcher/config"
	"cooool-blog-api/internal/watcher/options"
	log "cooool-blog-api/pkg/rollinglog"
	"cooool-blog-api/pkg/shutdown"
	"cooool-blog-api/pkg/shutdown/shutdownmanagers/posixsignal"
)

type watcherServer struct {
	gs             *shutdown.GracefulShutdown
	cron           *watchJob
	redisOptions   *genericoptions.RedisOptions
	mysqlOptions   *genericoptions.MySQLOptions
	watcherOptions *options.WatcherOptions
}

// preparedGenericAPIServer is a private wrapper that enforces a call of BeforeRun() before Run can be invoked.
type preparedWatcherServer struct {
	*watcherServer
}

func createWatcherServer(cfg *config.Config) *watcherServer {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	server := &watcherServer{
		gs:             gs,
		redisOptions:   cfg.RedisOptions,
		mysqlOptions:   cfg.MySQLOptions,
		watcherOptions: cfg.WatcherOptions,
	}

	return server
}

// BeforeRun prepares the server to run, by setting up the server instance.
func (s *watcherServer) BeforeRun() preparedWatcherServer {
	mysqlStore, err := mysql.GetMySQLFactoryOr(s.mysqlOptions)
	if err != nil {
		panic(err)
	}

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		return mysqlStore.Close()
	}))

	s.cron = newWatchJob(s.redisOptions, s.watcherOptions).addWatchers()

	return preparedWatcherServer{s}
}

func (s preparedWatcherServer) Run() error {
	stopCh := make(chan struct{})
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		// wait for running jobs to complete.
		ctx := s.cron.Stop()
		select {
		case <-ctx.Done():
			log.Info("cron jobs stopped.")
		case <-time.After(3 * time.Minute):
			log.Error("context was not done after 3 minutes.")
		}

		return nil
	}))

	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	log.Info("star to run cron jobs.")
	s.cron.Start()

	// blocking here via channel to prevents the process exit.
	<-stopCh

	return nil
}
