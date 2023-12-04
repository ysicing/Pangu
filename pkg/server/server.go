package server

import (
	"net/http"
	"os"

	"gitea.ysicing.net/cloud/pangu/common"
	"github.com/ergoapi/util/exgin"
	"github.com/ergoapi/util/exhttp"
	"github.com/sirupsen/logrus"
)

func Serve() error {
	g := exgin.Init(&exgin.Config{
		Debug:   true,
		Gops:    true,
		Pprof:   true,
		Metrics: true,
	})
	g.Use(exgin.ExLog(), exgin.ExRecovery(), exgin.Translations())
	addr := "0.0.0.0:65001"
	srv := &http.Server{
		Addr:    addr,
		Handler: g,
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
