/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	onlyOenSignalHandle = make(chan struct{})
	shutdownHandle      chan os.Signal
	shutdownSignals     = []os.Signal{os.Interrupt, syscall.SIGTERM}
)

func SetupSignaHandle() <-chan struct{} {
	close(onlyOenSignalHandle)

	shutdownHandle = make(chan os.Signal, 2)
	stop := make(chan struct{})

	signal.Notify(shutdownHandle, shutdownSignals...)

	go func() {
		<-shutdownHandle
		close(stop)
		<-shutdownHandle
		os.Exit(1)
	}()
	return stop
}
