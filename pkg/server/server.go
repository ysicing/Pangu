package server

import (
	"context"
	"net/http"
	"os"

	"github.com/ergoapi/util/exgin"
	"github.com/ergoapi/util/exhttp"
	"github.com/sirupsen/logrus"
)

func Serve(ctx context.Context) error {
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
		exhttp.SetupGracefulStop(srv)
		logrus.Info("server exited.")
	}()
	logrus.Infof("http listen to %v, pid is %v", addr, os.Getpid())
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("Failed to start http server, error: %s", err)
		return err
	}
	return nil
}
