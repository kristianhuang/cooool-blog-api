/*
 * Copyright 2021 SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"strings"
	"time"
)

type GenericAPIServer struct {
	middlewares []string
	mode        string
	// SecureServingInfo holds configuration of the TLS server.
	SecureServingInfo *SecureServingInfo

	// InsecureServingInfo holds configuration of the insecure HTTP server.
	InsecureServingInfo *InsecureServingInfo
	ShutdownTimeout     time.Duration

	*gin.Engine
	health          bool
	enableMetrics   bool
	enableProfiling bool

	insecureServer, secureServer *http.Server
}

func (s *GenericAPIServer) InstallAPIs() {
	// TODO 健康检测
	if s.health {
		s.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{})
		})
	}

	// 开启 gin 监控
	if s.enableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}

	// 开启性能分析
	if s.enableProfiling {
		pprof.Register(s.Engine)
	}

	// TODO 版本管理功能
}

func initGenericAPIServer(s *GenericAPIServer) {
	s.Setup()
	s.InstallMiddlewares()
	s.InstallAPIs()
}

func (s *GenericAPIServer) InstallMiddlewares() {
	// this is global middleware
}

func (s *GenericAPIServer) Setup() {
	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		// TODO 换成日志包输出
		fmt.Println(httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *GenericAPIServer) Run() error {
	s.insecureServer = &http.Server{
		Addr:    s.InsecureServingInfo.Address,
		Handler: s,
	}

	// https
	s.secureServer = &http.Server{
		Addr:    s.SecureServingInfo.Address(),
		Handler: s,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
			return err
		}

		return nil
	})

	// https
	// eg.Go(func() error {
	//
	//
	// 	return nil
	// })

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if s.health {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (s *GenericAPIServer) ping(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", s.InsecureServingInfo.Address)
	if strings.Contains(s.InsecureServingInfo.Address, "0.0.0.0") {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(s.InsecureServingInfo.Address, ":")[1])
	}

	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Println("The router has been deployed success")
			resp.Body.Close()
			return nil
		}

		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			log.Println("can not ping http server.")
		default:
		}
	}
}

// func WriteResponse(c *gin.Context, err error, data interface{}) {
// 	if err != nil {
// 		log.Errorf("%#+v", err)
// 		coder := errors.ParseCoder(err)
// 		c.JSON(coder.HTTPStatus(), ErrResponse{
// 			Code:      coder.Code(),
// 			Message:   coder.String(),
// 			Reference: coder.Reference(),
// 		})
//
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, data)
// }
