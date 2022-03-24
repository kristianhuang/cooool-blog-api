/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package authzserver

import (
	"context"

	"cooool-blog-api/internal/authzserver/analytics"
	"cooool-blog-api/internal/authzserver/config"
	"cooool-blog-api/internal/authzserver/load"
	"cooool-blog-api/internal/authzserver/load/cache"
	"cooool-blog-api/internal/authzserver/store/apiserver"
	genericoptions "cooool-blog-api/internal/pkg/options"
	genericapiserver "cooool-blog-api/internal/pkg/server"
	"cooool-blog-api/pkg/errors"
	log "cooool-blog-api/pkg/rollinglog"
	"cooool-blog-api/pkg/shutdown"
	"cooool-blog-api/pkg/shutdown/shutdownmanagers/posixsignal"
	"cooool-blog-api/pkg/storage"
)

const RedisKeyPrefix = "analytics-"

type authzServer struct {
	gs               *shutdown.GracefulShutdown
	rpcServer        string
	clientCA         string
	redisOptions     *genericoptions.RedisOptions
	genericAPIServer *genericapiserver.GenericAPIServer
	analyticsOptions *analytics.AnalyticsOptions
	redisCancelFunc  context.CancelFunc
}

type preparedAuthzServer struct {
	*authzServer
}

// func createAuthzServer(cfg *config.Config) (*authzServer, error) {.
func createAuthzServer(cfg *config.Config) (*authzServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &authzServer{
		gs:               gs,
		redisOptions:     cfg.RedisOptions,
		analyticsOptions: cfg.AnalyticsOptions,
		rpcServer:        cfg.RPCServer,
		clientCA:         cfg.ClientCA,
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *authzServer) BeforeRun() preparedAuthzServer {
	_ = s.initialize()

	initRouter(s.genericAPIServer.Engine)

	return preparedAuthzServer{s}
}

// Run start to run AuthzServer.
func (s preparedAuthzServer) Run() error {
	stopCh := make(chan struct{})

	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	// nolint: errcheck
	go s.genericAPIServer.Run()

	// in order to ensure that the reported data is not lost,
	// please ensure the following graceful shutdown sequence
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		s.genericAPIServer.Close()
		if s.analyticsOptions.Enable {
			analytics.GetAnalytics().Stop()
		}
		s.redisCancelFunc()

		return nil
	}))

	// blocking here via channel to prevents the process exit.
	<-stopCh

	return nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.ServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

func (s *authzServer) buildStorageConfig() *storage.Config {
	return &storage.Config{
		Host:                  s.redisOptions.Host,
		Port:                  s.redisOptions.Port,
		Addrs:                 s.redisOptions.Addrs,
		MasterName:            s.redisOptions.MasterName,
		Username:              s.redisOptions.Username,
		Password:              s.redisOptions.Password,
		Database:              s.redisOptions.Database,
		MaxIdle:               s.redisOptions.MaxIdle,
		MaxActive:             s.redisOptions.MaxActive,
		Timeout:               s.redisOptions.Timeout,
		EnableCluster:         s.redisOptions.EnableCluster,
		UseSSL:                s.redisOptions.UseSSL,
		SSLInsecureSkipVerify: s.redisOptions.SSLInsecureSkipVerify,
	}
}

func (s *authzServer) initialize() error {
	ctx, cancel := context.WithCancel(context.Background())
	s.redisCancelFunc = cancel

	// keep redis connected
	go storage.ConnectToRedis(ctx, s.buildStorageConfig())

	// cron to reload all secrets and policies from apiserver
	cacheIns, err := cache.GetCacheInsOr(apiserver.GetAPIServerFactoryOrDie(s.rpcServer, s.clientCA))
	if err != nil {
		return errors.Wrap(err, "get cache instance failed")
	}

	load.NewLoader(ctx, cacheIns).Start()

	// start analytics service
	if s.analyticsOptions.Enable {
		analyticsStore := storage.RedisCluster{KeyPrefix: RedisKeyPrefix}
		analyticsIns := analytics.NewAnalytics(s.analyticsOptions, &analyticsStore)
		analyticsIns.Start()
	}

	return nil
}
