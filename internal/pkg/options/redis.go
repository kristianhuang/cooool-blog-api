/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import "github.com/spf13/pflag"

type RedisOptions struct {
	Host                  string   `json:"host"`
	Port                  int      `json:"port"`
	Address               []string `json:"address"`
	Username              string   `json:"username"`
	Password              string   `json:"password"`
	Index                 int      `json:"index"`
	MasterName            string   `json:"master_name"`
	MaxIdle               int      `json:"max_idle"`
	MaxActive             int      `json:"max_active"`
	Timeout               int      `json:"timeout"`
	EnableCluster         bool     `json:"enable_cluster"`
	UseSSL                bool     `json:"use_ssl"`
	SSLInsecureSkipVerify bool     `json:"ssl_insecure_skip_verify"`
}

func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:                  "127.0.0.1",
		Port:                  6379,
		Address:               []string{},
		Username:              "",
		Password:              "",
		Index:                 0,
		MasterName:            "",
		MaxIdle:               2000,
		MaxActive:             4000,
		Timeout:               0,
		EnableCluster:         false,
		UseSSL:                false,
		SSLInsecureSkipVerify: false,
	}
}

func (o *RedisOptions) Validate() []error {
	return []error{}
}

func (o *RedisOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "redis.host", o.Host, "Redis host")
	fs.IntVar(&o.Port, "redis.port", o.Port, "Redis port")
	fs.StringSliceVar(&o.Address, "redis address", o.Address, "Redis address")
	fs.StringVar(&o.Username, "redis.username", o.Username, "Redis username")
	fs.StringVar(&o.Password, "redis.password", o.Password, "Redis password")
	fs.IntVar(&o.Index, "redis.index", o.Index, "Redis index")
	fs.StringVar(&o.MasterName, "redis.master-name", o.MasterName, "The name of master redis instance.")
	fs.IntVar(&o.MaxIdle, "redis.optimisation-max-idle", o.MaxIdle, ""+
		"This setting will configure how many connections are maintained in the pool when idle (no traffic). "+
		"Set the --redis.optimisation-max-active to something large, we usually leave it at around 2000 for "+
		"HA deployments.")
	fs.IntVar(&o.MaxActive, "redis.optimisation-max-active", o.MaxActive, ""+
		"In order to not over commit connections to the Redis server, we may limit the total "+
		"number of active connections to Redis. We recommend for production use to set this to around 4000.")
	fs.IntVar(&o.Timeout, "redis.timeout", o.Timeout, "Timeout (in seconds) when connecting to redis service.")
	fs.BoolVar(&o.EnableCluster, "redis.enable-cluster", o.EnableCluster, ""+
		"If you are using Redis cluster, enable it here to enable the slots mode.")
	fs.BoolVar(&o.UseSSL, "redis.use-ssl", o.UseSSL, ""+
		"If set, IAM will assume the connection to Redis is encrypted. "+
		"(use with Redis providers that support in-transit encryption).")
	fs.BoolVar(&o.SSLInsecureSkipVerify, "redis.ssl-insecure-skip-verify", o.SSLInsecureSkipVerify, ""+
		"Allows usage of self-signed certificates when connecting to an encrypted Redis database.")

}
