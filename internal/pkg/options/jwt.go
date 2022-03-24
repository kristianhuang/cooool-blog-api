/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"fmt"
	"time"

	"cooool-blog-api/internal/pkg/server"

	"github.com/asaskevich/govalidator"
	"github.com/spf13/pflag"
)

type JwtOptions struct {
	Realm      string        `json:"realm" mapstructure:"realm"`
	Key        string        `json:"key" mapstructure:"key"`
	Timeout    time.Duration `json:"timeout" mapstructure:"timeout"`
	MaxRefresh time.Duration `json:"max_refresh" mapstructure:"max_refresh"`
}

// NewJwtOptions creates a JwtOptions object with default params.
func NewJwtOptions() *JwtOptions {
	defaults := server.NewConfig()
	return &JwtOptions{
		Realm:      defaults.Jwt.Realm,
		Key:        defaults.Jwt.Key,
		Timeout:    defaults.Jwt.Timeout,
		MaxRefresh: defaults.Jwt.MaxRefresh,
	}
}

func (o *JwtOptions) ApplyTo(c *server.Config) error {
	c.Jwt = &server.JwtInfo{
		Realm:      o.Realm,
		Key:        o.Key,
		Timeout:    o.Timeout,
		MaxRefresh: o.MaxRefresh,
	}

	return nil
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *JwtOptions) Validate() []error {
	var errs []error

	if !govalidator.StringLength(o.Key, "6", "32") {
		errs = append(errs, fmt.Errorf("--secret-key must larger than 5 and little than 33"))
	}

	return errs
}

func (o *JwtOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&o.Realm, "jwt.realm", o.Realm, "Realm name to display to the user.")
	fs.StringVar(&o.Key, "jwt.key", o.Key, "Private key used to sign jwt token.")
	fs.DurationVar(&o.Timeout, "jwt.timeout", o.Timeout, "JWT token timeout.")

	fs.DurationVar(&o.MaxRefresh, "jwt.max-refresh", o.MaxRefresh, ""+
		"This field allows clients to refresh their token until MaxRefresh has passed.")
}
