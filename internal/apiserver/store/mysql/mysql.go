/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mysql

import (
	"fmt"
	"sync"

	"cooool-blog-api/internal/apiserver/store"
	"cooool-blog-api/internal/pkg/logger"
	"cooool-blog-api/internal/pkg/model"
	genericoptions "cooool-blog-api/internal/pkg/options"
	"cooool-blog-api/pkg/db"

	"gorm.io/gorm"
)

type dataStore struct {
	db *gorm.DB
}

func (ds *dataStore) AdminUser() store.AdminUserStore {
	return newAdminUser(ds)
}

func (ds *dataStore) Policies() store.PolicyStore {
	return newPolicy(ds)
}

func (ds *dataStore) PolicyAudits() store.PolicyAuditStore {
	return newPolicyAudits(ds)
}

func (ds *dataStore) Secrets() store.SecretStore {
	return newSecrets(ds)
}

func (ds *dataStore) Close() error {
	db, err := ds.db.DB()

	if err != nil {
		return err
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (store.Factory, error) {
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

	if err := db.AutoMigrate(&model.Policy{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Secret{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Secret{}); err != nil {
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
