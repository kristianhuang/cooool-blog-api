/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"blog-go/internal/apiserver/options"
	"blog-go/pkg/app"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	// fmt.Println(opts)

	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	app.NewApp("api-server", "blog-server-go",
		app.WithFlags(options.NewOptions()),
		app.WithRunFunc(func(basename string) error {
			fmt.Println(basename)
			return nil
		})).Run()
}
