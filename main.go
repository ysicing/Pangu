package main

import (
	"os"

	"gitea.ysicing.net/cloud/pangu/pkg/server"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
}

func main() {
	if err := server.Serve(); err != nil {
		logrus.Fatalf("run serve: %v", err)
	}
}
