// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"strings"

	"gitea.ysicing.net/cloud/pangu/common"
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/zos"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pangu",
		Short: "pangu sdk by ysicing",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(false)
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&common.CfgFile, "config", "", "config file (default is /conf/example.yml)")
	rootCmd.PersistentFlags().BoolVar(&common.Debug, "debug", false, "enable debug logging")
	rootCmd.AddCommand(webCommand())
}

func initConfig() {
	if common.CfgFile == "" {
		common.CfgFile = common.GetDefaultPath()
		if zos.IsMacOS() {
			common.CfgFile = "./pangu.yaml"
		}
	}
	viper.SetConfigFile(common.CfgFile)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PGU")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	if err := viper.ReadInConfig(); err == nil {
		logrus.Debugf("Using config file: %v", color.SGreen(viper.ConfigFileUsed()))
	}
	// reload
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Debugf("config changed: %v", color.SGreen(in.Name))
	})
	if common.Debug {
		viper.Set("debug", true)
	}
}
