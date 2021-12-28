/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

type CertKey struct {
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
}

type GeneratableKeyCert struct {
	CertKey  CertKey `json:"cert_key"`
	CertDir  string  `json:"cert_dir"`
	PairName string  `json:"pair_name"`
}

type SecureServingOptions struct {
	BindAddress string             `json:"bind_address"`
	BindPort    int                `json:"bind_port"`
	Required    bool               `json:"required"`
	ServerCert  GeneratableKeyCert `json:"server_cert"`
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8080,
		Required:    true,
		ServerCert: GeneratableKeyCert{
			PairName: "blog-api",
			CertDir:  "/www/blog-api",
		},
	}
}
