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
	"blog-api/internal/pkg/logger"
	"blog-api/internal/pkg/model"
	genericoptions "blog-api/internal/pkg/options"
	"blog-api/pkg/db"
	"gorm.io/gorm"
)

type dataStore struct {
	db *gorm.DB
}

func (s *dataStore) AdminUser() store.AdminUserStore {
	return newAdminUser(s)
}

func (s *dataStore) Policy() store.PolicyStore {
	return newPolicy(s)
}

func (s *dataStore) Close() error {
	ds, err := s.db.DB()

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
			Logger:                logger.NewLogger(opts.LogLevel),
		}

		dbIns, err = db.New(options)
		mysqlFactory = &dataStore{dbIns}

		err = MigrateDatabase(dbIns)
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get db store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

func cleanDatabases(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&model.AdminUser{}); err != nil {
		return err
	}

	return nil
}

func MigrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.AdminUser{}); err != nil {
		return err
	}

	return nil
}

func ResetDatabase(db *gorm.DB) error {
	if err := cleanDatabases(db); err != nil {
		return err
	}

	if err := MigrateDatabase(db); err != nil {
		return err
	}

	return nil
}
