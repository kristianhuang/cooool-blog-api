/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package analytics

import (
	"sync"
	"sync/atomic"
	"time"

	log "cooool-blog-api/pkg/rollinglog"

	"gopkg.in/vmihailenco/msgpack.v2"

	"cooool-blog-api/pkg/storage"
)

const (
	analyticsKeyName = "blog-system-analytics"

	recordsBufferForcedFlushInterval = 1 * time.Second
)

type AnalyticsRecord struct {
	Timestamp  int64     `json:"timestamp"`
	Username   string    `json:"username"`
	Effect     string    `json:"effect"`
	Conclusion string    `json:"conclusion"`
	Request    string    `json:"request"`
	Policies   string    `json:"policies"`
	Deciders   string    `json:"deciders"`
	ExpireAt   time.Time `json:"expire_at" bson:"expire_at"`
}

var analytics *Analytics

func (a *AnalyticsRecord) SetExpiry(expiresInSeconds int64) {
	expiry := time.Duration(expiresInSeconds) * time.Second
	if expiry == 0 {
		// Expiry is set to 100 years
		expiry = 24 * 365 * 100 * time.Hour
	}

	t := time.Now()
	t2 := t.Add(expiry)
	a.ExpireAt = t2
}

type Analytics struct {
	store                      storage.AnalyticsHandler
	poolSize                   int
	recordsChan                chan *AnalyticsRecord
	workerBufferSize           uint64
	recordsBufferFlushInterval uint64
	shouldStop                 uint32
	poolWg                     sync.WaitGroup
}

// NewAnalytics returns a new analytics instance.
func NewAnalytics(options *AnalyticsOptions, store storage.AnalyticsHandler) *Analytics {
	ps := options.PoolSize
	recordsBufferSize := options.RecordsBufferSize
	workerBufferSize := recordsBufferSize / uint64(ps)
	recordsChan := make(chan *AnalyticsRecord, recordsBufferSize)

	analytics = &Analytics{
		store:                      store,
		poolSize:                   ps,
		recordsChan:                recordsChan,
		workerBufferSize:           workerBufferSize,
		recordsBufferFlushInterval: options.FlushInterval,
	}

	return analytics
}

// GetAnalytics returns the existed analytics instance.
// Need to initialize `analytics` instance before calling GetAnalytics.
func GetAnalytics() *Analytics {
	return analytics
}

// Start start the analytics service.
func (a *Analytics) Start() {
	a.store.Connect()

	atomic.SwapUint32(&a.shouldStop, 0)

	for i := 0; i < a.poolSize; i++ {
		a.poolWg.Add(1)
		go a.recordWorker()
	}
}

// Stop stop the analytics service.
func (a *Analytics) Stop() {
	// flag to stop sending records into channel
	atomic.SwapUint32(&a.shouldStop, 1)

	// close channel to stop workers
	close(a.recordsChan)

	// wait for all workers to be done
	a.poolWg.Wait()
}

// RecordHit will store an AnalyticsRecord in Redis.
func (a *Analytics) RecordHit(record *AnalyticsRecord) error {
	// check if we should stop sending records 1st
	if atomic.LoadUint32(&a.shouldStop) > 0 {
		return nil
	}

	// just send record to channel consumed by pool of workers
	// leave all data crunching and Redis I/O work for pool workers
	a.recordsChan <- record

	return nil
}

func (a *Analytics) recordWorker() {
	defer a.poolWg.Done()

	// this is buffer to send one pipelined command to redis
	// use r.recordsBufferSize as cap to reduce slice re-allocations
	recordsBuffer := make([][]byte, 0, a.workerBufferSize)

	// read records from channel and process
	lastSentTS := time.Now()

	for {
		var readyToSend bool

		select {
		case record, ok := <-a.recordsChan:
			if !ok {
				a.store.AppendToSetPipelined(analyticsKeyName, recordsBuffer)
				return
			}

			// we have new record - prepare it and add to buffer

			if encoded, err := msgpack.Marshal(record); err != nil {
				log.Errorf("Error encoding analytics data: %s", err.Error())
			} else {
				recordsBuffer = append(recordsBuffer, encoded)
			}

			readyToSend = uint64(len(recordsBuffer)) == a.workerBufferSize

		case <-time.After(time.Duration(a.recordsBufferFlushInterval) * time.Millisecond):
			readyToSend = true
		}

		if len(recordsBuffer) > 0 && (readyToSend || time.Since(lastSentTS) >= recordsBufferForcedFlushInterval) {
			a.store.AppendToSetPipelined(analyticsKeyName, recordsBuffer)
			recordsBuffer = recordsBuffer[:0]
			lastSentTS = time.Now()
		}
	}
}

func DurationToMillisecond(d time.Duration) float64 {
	return float64(d) / 1e6
}
