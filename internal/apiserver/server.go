/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	genericoptions "blog-go/internal/pkg/options"
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

// func buildGenericConf(c *config.Config) (genericConf *genericap) {
//
// }
