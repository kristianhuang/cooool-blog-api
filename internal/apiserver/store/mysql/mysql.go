/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"fmt"
	"sync"

	"blog-api/internal/apiserver/store"
	"blog-api/internal/pkg/gormutil"
	genericoptions "blog-api/internal/pkg/options"
	"blog-api/pkg/db"
	"gorm.io/gorm"
)

type dataStore struct {
	*gorm.DB
}

func (s *dataStore) AdminUser() store.AdminUserStore {
	return newAdminUser(s.DB)
}

func (s *dataStore) Close() error {
	ds, err := s.DB.DB()

	if err != nil {
		return err
	}

	return ds.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMysqlFactory(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("faailed to get db store factory")
	}
	var (
		err   error
		dbIns *gorm.DB
	)
	once.Do(func() {
		options := &db.MysqlOptions{
			Host:                  opts.Host,
			Port:                  opts.Port,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                gormutil.NewLogger(opts.LogLevel),
		}

		dbIns, err = db.New(options)

		mysqlFactory = &dataStore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get db store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}
