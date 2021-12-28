/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import (
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
)

type Conf struct {
	Mode        string
	Middlewares []string
	Health      bool
	// 启用配置文件
	EnableProfiling bool
	// 启用统计
	EnableMetrics bool
}

type SecureServingInfo struct {
	BindAddress string
	BindPort    int
}

type InsecureServingInfo struct {
	Address string
}

func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

func NewConf() *Conf {
	return &Conf{
		Mode:            gin.ReleaseMode,
		Health:          true,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
	}
}

type CompletedConf struct {
	*Conf
}

func (c *Conf) Complete() CompletedConf {
	return CompletedConf{c}
}

// func (c CompletedConf) New() (*Gen) {
//
// }
