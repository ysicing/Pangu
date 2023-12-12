// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"gitea.ysicing.net/cloud/pangu/common"
	_ "gitea.ysicing.net/cloud/pangu/docs"
	"gitea.ysicing.net/cloud/pangu/internal/cache"
	"gitea.ysicing.net/cloud/pangu/internal/cron"
	"gitea.ysicing.net/cloud/pangu/internal/db"
	"gitea.ysicing.net/cloud/pangu/internal/routes"
	_ "gitea.ysicing.net/cloud/pangu/internal/routes/v1/config"
	_ "gitea.ysicing.net/cloud/pangu/internal/routes/v1/custom"
	"gitea.ysicing.net/cloud/pangu/internal/service/config"
	"gitea.ysicing.net/cloud/pangu/internal/service/user"
	"gitea.ysicing.net/cloud/pangu/pkg/util"
	"github.com/cockroachdb/errors"
	"github.com/ergoapi/util/exgin"
	"github.com/ergoapi/util/exhttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve() error {
	if err := db.SetDB(); err != nil {
		return err
	}
	if err := cache.SetCache(); err != nil {
		return err
	}
	if err := InitData(); err != nil {
		return errors.Errorf("初始化数据异常: %v", err)
	}
	defer cron.Cron.Stop()
	cron.Cron.Start()
	metricsPath := util.GetKeyFromYaml("server.metrics.path", "/ops/metrics")
	docsPath := strings.TrimSuffix(util.GetKeyFromYaml("server.docs", "/docs"), "/")
	g := exgin.Init(&exgin.Config{
		Debug:       viper.GetBool("debug"),
		Gops:        util.GetStatusFromYaml("server.gops"),
		Pprof:       util.GetStatusFromYaml("server.pprof"),
		MetricsPath: metricsPath,
		Metrics:     true,
	})
	g.Use(exgin.ExLog(metricsPath, docsPath), exgin.ExRecovery(), exgin.Translations())
	g.GET(fmt.Sprintf("%s/*any", docsPath), ginSwagger.WrapHandler(swaggerfiles.Handler))
	g.NoMethod(func(ctx *gin.Context) {
		ctx.String(http.StatusForbidden, "Forbidden")
	})
	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusForbidden, "Forbidden")
	})
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

func InitData() error {
	if config.Init() {
		logrus.Info("initialized")
		return nil
	}
	if err := user.Init(); err != nil {
		return err
	}
	return config.InitDone()
}
