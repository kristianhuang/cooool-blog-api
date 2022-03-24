/*
 * Copyright 2021 Kris Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package watcher

import (
	"context"
	"fmt"
	"time"

	genericoptions "cooool-blog-api/internal/pkg/options"
	"cooool-blog-api/internal/watcher/options"
	"cooool-blog-api/internal/watcher/watcher"
	log "cooool-blog-api/pkg/rollinglog"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/marmotedu/iam/pkg/log/cronlog"
	"github.com/robfig/cron/v3"
)

type watchJob struct {
	*cron.Cron
	config *options.WatcherOptions
	rs     *redsync.Redsync
}

func newWatchJob(redisOptions *genericoptions.RedisOptions, watcherOptions *options.WatcherOptions) *watchJob {
	logger := cronlog.NewLogger(log.SugaredLogger())

	client := goredislib.NewClient(&goredislib.Options{
		Addr:     fmt.Sprintf("%s:%d", redisOptions.Host, redisOptions.Port),
		Username: redisOptions.Username,
		Password: redisOptions.Password,
	})

	rs := redsync.New(goredis.NewPool(client))

	cron := cron.New(
		cron.WithSeconds(),
		cron.WithChain(cron.SkipIfStillRunning(logger), cron.Recover(logger)),
	)

	return &watchJob{
		Cron:   cron,
		config: watcherOptions,
		rs:     rs,
	}
}

func (w *watchJob) addWatchers() *watchJob {
	for name, watcher := range watcher.ListWatchers() {
		// log with `{"watcher": "counter"}` key-value to distinguish which watcher the log comes from.
		// nolint: golint,staticcheck
		ctx := context.WithValue(context.Background(), log.KeyWatcherName, name)

		if err := watcher.Init(ctx, w.rs.NewMutex(name, redsync.WithExpiry(2*time.Hour)), w.config); err != nil {
			log.Panicf("construct watcher %s failed: %s", name, err.Error())
		}

		_, _ = w.AddJob(watcher.Spec(), watcher)
	}

	return w
}
