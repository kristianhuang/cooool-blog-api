/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-go/internal/apiserver/config"
	genericoptions "blog-go/internal/pkg/options"
	"blog-go/internal/pkg/server"
	"blog-go/internal/pkg/shutdown"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	redisOpts        *genericoptions.RedisOptions
	genericAPIServer *genericoptions.APIServerOptions
}

// func createApiServer(c *config.Config) (*apiServer, error) {
// 	// TODO 优雅关闭
//
// }

type perparedAPIServer struct {
	*apiServer
}

type ExtraConf struct {
	Addr       string
	MaxMsgSize int
	mysqlOpts  *genericoptions.MySQLOptions
}

func createAPIServer(conf *config.Config) (*apiServer, error) {
	// genericConf, err := buildGenericConf(conf)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// // genericServer, err := genericConf.Complete()
	// // return &apiServer{
	// // 	redisOpts:        conf.RedisOptions,
	// // 	genericAPIServer: genericConf,
	// // }, nil

	return nil, nil
}

func buildGenericConf(conf *config.Config) (genericConf *server.Conf, err error) {
	genericConf = server.NewConf()
	if err = conf.APISServerOptions.ApplyTo(genericConf); err != nil {
		return
	}

	return
}

func buildExtraConf(conf *config.Config) (*ExtraConf, error) {
	return &ExtraConf{
		mysqlOpts: conf.MySQLOptions,
	}, nil
}
