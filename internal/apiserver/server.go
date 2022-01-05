/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-go/internal/apiserver/config"
	genericapiserver "blog-go/internal/pkg/server"
	"blog-go/internal/pkg/shutdown"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

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

func createAPIServer(config *config.Config) (*apiServer, error) {
	// gs := shutdown.New()
	// gs.AddShutdownCallback()
	genericAPIServerConfig, err := buildGenericConfig(config)
	if err != nil {
		return nil, err
	}

	genericAPIServer, err := genericAPIServerConfig.Complete().NewGenericAPIServer()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		genericAPIServer: genericAPIServer,
	}

	return server, nil
}

func (s *apiServer) BeforeRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {

	return s.genericAPIServer.Run()
}
