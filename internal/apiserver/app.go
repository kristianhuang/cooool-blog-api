/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package apiserver

import (
	"blog-go/internal/apiserver/options"
	"blog-go/pkg/app"
	"fmt"
	"github.com/fatih/color"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.
Find more iam-apiserver information at:
    https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-apiserver.md`

func NewApp(baseName string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(baseName, "API Server", app.WithFlags(opts), app.WithLong(commandDesc), app.WithRunFunc(runFunc(opts)))

	return application
}

func runFunc(opts *options.Options) app.RunFunc {

	return func(basename string) error {
		fmt.Println(color.BlueString("i am run"))
		return nil
	}
}
