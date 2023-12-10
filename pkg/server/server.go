// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package server

import (
	"net/http"
	"os"
	"time"

	"gitea.ysicing.net/cloud/pangu/common"
	_ "gitea.ysicing.net/cloud/pangu/docs"
	"gitea.ysicing.net/cloud/pangu/internal/routes"
	_ "gitea.ysicing.net/cloud/pangu/internal/routes/v1/config"
	_ "gitea.ysicing.net/cloud/pangu/internal/routes/v1/custom"
	"github.com/ergoapi/util/exgin"
	"github.com/ergoapi/util/exhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve() error {
	g := exgin.Init(&exgin.Config{
		Debug:   viper.GetBool("debug"),
		Gops:    true,
		Pprof:   true,
		Metrics: true,
	})
	g.Use(exgin.ExLog(), exgin.ExRecovery(), exgin.Translations())
	g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.SetupRoutes(g)
	addr := "0.0.0.0:65001"
	srv := &http.Server{
		Addr:              addr,
		Handler:           g,
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		logrus.Infof("Version: %s, http listen to %v, pid is %v", common.GetVersion(), addr, os.Getpid())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Panicf("Failed to start http server, error: %s", err)
		}
	}()
	exhttp.SetupGracefulStop(srv)
	return nil
}
