/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"fmt"
	"net"
	"strconv"

	"blog-go/internal/pkg/server"
	"github.com/spf13/pflag"
)

type InsecureServingOptions struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewInsecureServingOptions() *InsecureServingOptions {
	return &InsecureServingOptions{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func (o *InsecureServingOptions) ApplyTo(c *server.Config) error {
	c.InsecureServing = &server.InsecureServingInfo{
		Host: net.JoinHostPort(o.Host, strconv.Itoa(o.Port)),
	}

	return nil
}

func (o *InsecureServingOptions) Validate() []error {
	var errors []error

	if o.Port < 0 || o.Port > 65535 {
		errors = append(errors, fmt.Errorf("--insecure.port %v must be between 0 and 65535", o.Port))
	}

	return errors
}

func (o InsecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "insecure.host", o.Host, "API server host")
	fs.IntVar(&o.Port, "insecure.port", o.Port, "API server port")
}
