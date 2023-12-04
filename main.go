package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"gitea.ysicing.net/cloud/pangu/pkg/server"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-ctx.Done()
		stop()
	}()

	if err := server.Serve(ctx); err != nil {
		logrus.Fatalf("run serve: %v", err)
	}
}
