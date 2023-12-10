// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package db

import (
	"strings"
	"time"

	"gitea.ysicing.net/cloud/pangu/common"
	"gitea.ysicing.net/cloud/pangu/pkg/util"
	"github.com/cockroachdb/errors"
	"github.com/ergoapi/util/log/glog"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/prometheus"
)

var DB *gorm.DB

var Migrates []interface{}

func Migrate(obj interface{}) {
	Migrates = append(Migrates, obj)
}

func SetDB() error {
	var err error
	dsn := util.GetKeyFromYaml("db.dsn", "file:/tmp/pangu.db?cache=shared&mode=rwc")

	dbcfg := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: &glog.DefaultGLogger,
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation("Asia/Shanghai")
			return time.Now().In(loc).Truncate(time.Second)
		},
	}
	logrus.Infof("load db config")
	if strings.Contains(dsn, "@") {
		DB, err = gorm.Open(mysql.Open(dsn), dbcfg)
	} else {
		DB, err = gorm.Open(sqlite.Open(dsn), dbcfg)
	}
	if err != nil {
		return errors.Errorf("连接数据库异常: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return errors.Errorf("获取数据库连接异常: %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(util.GetKeyIntFromYaml("db.max_idle_conns", 10))
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(util.GetKeyIntFromYaml("db.max_open_conns", 100))
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Minute)

	if util.GetStatusFromYaml("db.metrics.enable") {
		dbname := util.GetKeyFromYaml("db.metrics.name", common.DBMetricsName)
		if err := DB.Use(prometheus.New(prometheus.Config{
			DBName: dbname,
		})); err != nil {
			return errors.Errorf("启用数据库监控扩展异常: %v", err)
		}
	}
	logrus.Debugf("table num: %d", len(Migrates))
	if err := DB.AutoMigrate(Migrates...); err != nil {
		return errors.Errorf("初始化数据库表结构异常: %v", err)
	}
	logrus.Infof("create db engine success...")
	return nil
}
