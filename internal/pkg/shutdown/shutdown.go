/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package shutdown

import "sync"

// Callback 停止运行时回调接口
type Callback interface {
	OnShutdown(string) error
}
type ShutdownFunc func(string) error

func (f ShutdownFunc) OnShutdown(manager string) error {
	return f(manager)
}

type Manager interface {
	GetName() string
	Start(gs GSInterface) error
	ShutdownStart() error
	ShutdownFinish() error
}

type ErrHandler interface {
	OnError(err error)
}

type ErrorFunc func(err error)

func (f ErrorFunc) OnError(err error) {
	f(err)
}

type GSInterface interface {
	StartShutdown(sm Manager)
	ReportErr(err error)
	AddShutdownCallback(sc Callback)
}

type GracefulShutdown struct {
	callbacks []Callback
	managers  []Manager
	errHandle ErrHandler
}

func New() *GracefulShutdown {
	return &GracefulShutdown{
		callbacks: make([]Callback, 0),
		managers:  make([]Manager, 0),
	}
}

func (gs *GracefulShutdown) Start() error {
	for _, manager := range gs.managers {
		if err := manager.Start(gs); err != nil {
			return err
		}
	}
	return nil
}

func (gs *GracefulShutdown) StartShutdown(sm Manager) {
	gs.ReportErr(sm.ShutdownStart())

	var wg sync.WaitGroup

	for _, callback := range gs.callbacks {
		wg.Add(1)
		go func(shutdownCallback Callback) {
			defer wg.Done()
			gs.ReportErr(shutdownCallback.OnShutdown(sm.GetName()))
		}(callback)
	}

	wg.Wait()

	gs.ReportErr(sm.ShutdownFinish())
}

func (gs *GracefulShutdown) ReportErr(err error) {
	if err != nil && gs.errHandle != nil {
		gs.errHandle.OnError(err)
	}
}

func (gs *GracefulShutdown) AddShutdownCallback(sc Callback) {
	gs.callbacks = append(gs.callbacks, sc)
}

func (gs *GracefulShutdown) AddManager(m Manager) {
	gs.managers = append(gs.managers, m)
}

func (gs *GracefulShutdown) SetErrHandle(errHandler ErrHandler) {
	gs.errHandle = errHandler
}
