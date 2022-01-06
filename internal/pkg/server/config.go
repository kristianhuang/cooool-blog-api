/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import (
	"log"
	"net"
	"path/filepath"
	"strconv"
	"strings"

	"blog-go/pkg/path/dir"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	RecommendedHomeDir   = ".blog"
	RecommendedEnvPrefix = "BLOG"
)

type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
	Mode            string
	Middlewares     []string
	Health          bool
	// 启用配置文件
	EnableProfiling bool
	// 启用统计
	EnableMetrics bool
}

type SecureServingInfo struct {
	Host string
	Port int
}

type InsecureServingInfo struct {
	Host string
	Port int
}

// Address return host:port.
func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
}

// Address return host:port.
func (i *InsecureServingInfo) Address() string {
	return net.JoinHostPort(i.Host, strconv.Itoa(i.Port))
}

func NewConfig() *Config {
	return &Config{
		Mode:            gin.DebugMode,
		Health:          false,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
		InsecureServing: &InsecureServingInfo{
			Host: "127.0.0.1",
			Port: 8080,
		},
	}
}

type CompletedConfig struct {
	*Config
}

// Complete return completed config
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

func (c CompletedConfig) CreateGenericServer() (*GenericServer, error) {
	s := &GenericServer{
		// SecureServingInfo:   c.SecureServing,
		InsecureServingInfo: c.InsecureServing,
		Engine:              gin.New(),
		mode:                c.Mode,
		health:              c.Health,
		enableProfiling:     c.EnableProfiling,
		enableMetrics:       c.EnableMetrics,
		middlewares:         c.Middlewares,
	}

	initGenericAPIServer(s)

	return s, nil
}

// LoadConfig read config file and ENV vars, If set.
func LoadConfig(config, defaultName string) {
	if config != "" {
		viper.SetConfigFile(config)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Join(dir.HomeDir(), RecommendedHomeDir))
		viper.AddConfigPath("/etc/blog")
		viper.SetConfigName(defaultName)
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(RecommendedEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", ")"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("WARNING: viper failed to discover and load the configuration file: %s", err.Error())
	}
}
