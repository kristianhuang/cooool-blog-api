/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package posixsignal

import (
	"os"
	"os/signal"
	"syscall"

	"cooool-blog-api/pkg/shutdown"
)

const Name = "PosixSignalManager"

type PosixSignalManager struct {
	signals []os.Signal
}

func NewPosixSignalManager(sig ...os.Signal) *PosixSignalManager {
	if len(sig) == 0 {
		sig = make([]os.Signal, 2)
		sig[0] = os.Interrupt
		sig[1] = syscall.SIGTERM
	}

	return &PosixSignalManager{
		signals: sig,
	}
}

func (m *PosixSignalManager) GetName() string {
	return Name
}

func (m *PosixSignalManager) Start(gs shutdown.GSInterface) error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, m.signals...)

		<-c

		gs.StartShutdown(m)
	}()

	return nil
}

func (m *PosixSignalManager) ShutdownStart() error {
	return nil
}

func (m *PosixSignalManager) ShutdownFinish() error {
	os.Exit(0)
	return nil
}
