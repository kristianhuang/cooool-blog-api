/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"blog-api/pkg/util/path/dir"
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
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Health          bool
	// 启用配置文件
	EnableProfiling bool
	// 启用统计
	EnableMetrics bool
}

type JwtInfo struct {
	// Defaults to "blog jwt"
	Realm string
	// Defaults to empty
	Key string
	// Defaults to one hour
	Timeout time.Duration
	// Defaults to one hour
	MaxRefresh time.Duration
}

type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey
}

type InsecureServingInfo struct {
	BindAddress string
	BindPort    int
}

// CertKey contains configuration items related to certificate.
type CertKey struct {
	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

// Address return host:port.
func (s *SecureServingInfo) Address() string {
	return ":" + strconv.Itoa(s.BindPort)
}

// Address return host:port.
func (i *InsecureServingInfo) Address() string {
	// return net.JoinHostPort(i.BindAddress, strconv.Itoa(i.BindPort))
	return ":" + strconv.Itoa(i.BindPort)
}

func NewConfig() *Config {
	return &Config{
		Mode:            gin.DebugMode,
		Health:          false,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
		InsecureServing: &InsecureServingInfo{
			BindAddress: "127.0.0.1",
			BindPort:    8080,
		},
		Jwt: &JwtInfo{
			Realm:      "api jwt",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
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

func (c CompletedConfig) New() (*GenericAPIServer, error) {
	s := &GenericAPIServer{
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
		viper.AddConfigPath("./config")
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
