/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-api/internal/apiserver/config"
	"blog-api/internal/apiserver/store"
	"blog-api/internal/apiserver/store/mysql"
	genericapiserver "blog-api/internal/pkg/server"
	"blog-api/pkg/shutdown"
	"blog-api/pkg/shutdown/shutdownmanagers/posixsignal"
)

type apiServer struct {
	gs            *shutdown.GracefulShutdown
	genericServer *genericapiserver.GenericServer
}

type preparedAPIServer struct {
	*apiServer
}

// 使用 apiserver 的配置项填充至 server 的配置项，用以满足启动 server 的必要条件。
func buildGenericConfig(conf *config.Config) (apiServerConfig *genericapiserver.Config, err error) {
	apiServerConfig = genericapiserver.NewConfig()
	if err = conf.ServerRunOptions.ApplyTo(apiServerConfig); err != nil {
		return
	}

	if err = conf.FeatureOptions.ApplyTo(apiServerConfig); err != nil {
		return
	}

	if err = conf.InsecureServingOptions.ApplyTo(apiServerConfig); err != nil {
		return
	}

	return
}

func createServer(config *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericServerConfig, err := buildGenericConfig(config)
	if err != nil {
		return nil, err
	}

	genericAPIServer, err := genericServerConfig.Complete().CreateGenericServer()
	if err != nil {
		return nil, err
	}

	// init mysql store.
	storeIns, _ := mysql.GetMysqlFactory(config.MySQLOptions)
	store.SetClient(storeIns)

	server := &apiServer{
		gs:            gs,
		genericServer: genericAPIServer,
	}

	return server, nil
}

func (s *apiServer) BeforeRun() preparedAPIServer {
	// init router
	initRouter(s.genericServer.Engine)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMysqlFactory(nil)
		if mysqlStore != nil {
			_ = mysqlStore.Close()
		}
		s.genericServer.Close()

		return nil
	}))

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	return s.genericServer.Run()
}
