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
	genericAPIServer *genericapiserver.APIServer
}

type preparedAPIServer struct {
	*apiServer
}

func buildGenericConfig(conf *config.Config) (genericConfig *genericapiserver.Config, err error) {
	genericConfig = genericapiserver.NewConfig()
	if err = conf.GenericServerRunOptions.ApplyTo(genericConfig); err != nil {
		return
	}

	if err = conf.FeatureOptions.ApplyTo(genericConfig); err != nil {
		return
	}

	if err = conf.InsecureServingOptions.ApplyTo(genericConfig); err != nil {
		return
	}

	return
}

func createAPIServer(conf *config.Config) (*apiServer, error) {
	// gs := shutdown.New()
	// gs.AddShutdownCallback()
	genericConfig, err := buildGenericConfig(conf)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		genericAPIServer: genericServer,
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
