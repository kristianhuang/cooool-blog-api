/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package options

import (
	"time"

	"blog-api/pkg/db"
	"github.com/spf13/pflag"
	"gorm.io/gorm"
)

type MySQLOptions struct {
	Host                  string        `json:"host" mapstructure:"host"`
	Port                  int           `json:"port" mapstructure:"port"`
	Username              string        `json:"username" mapstructure:"username"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max_idle_connections,omitempty" mapstructure:"max_idle_connections"`
	MaxOpenConnections    int           `json:"max_open_connections,omitempty" mapstructure:"max_open_connections"`
	MaxConnectionLifeTime time.Duration `json:"max_connection_life_time,omitempty" mapstructure:"max_connection_life_time"`
	LogLevel              int           `json:"log_level" mapstructure:"log_level"`
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:                  "127.0.0.1",
		Port:                  3306,
		Username:              "",
		Password:              "",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
		LogLevel:              1,
	}
}

func (o *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "db.host", o.Host, "MySQL host")
	fs.IntVar(&o.Port, "db.port", o.Port, "MySQL port")
	fs.StringVar(&o.Username, "db.username", o.Username, "MySQL username")
	fs.StringVar(&o.Password, "db.password", o.Password, "MySQL password")
	fs.StringVar(&o.Database, "db.database", o.Database, "MySQL database")
	fs.IntVar(&o.MaxIdleConnections, "db.max-idle-connections", o.MaxIdleConnections, "MySQL max-idle-connections")
	fs.IntVar(&o.MaxOpenConnections, "db.xax-open-connections", o.MaxOpenConnections, "MySQL xax-open-connections")
	fs.DurationVar(&o.MaxConnectionLifeTime, "db.max-connection-life-time", o.MaxConnectionLifeTime, "MySQL max-connection-life-time")
	fs.IntVar(&o.LogLevel, "db.log-mode", o.LogLevel, "GORM log level")
}

func (o *MySQLOptions) Validate() []error {
	return []error{}
}

func (o *MySQLOptions) NewClient() (*gorm.DB, error) {
	opts := &db.MysqlOptions{
		Host:                  o.Host,
		Port:                  o.Port,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
	}

	return db.New(opts)
}
