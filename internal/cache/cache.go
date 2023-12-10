// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cache

import (
	"context"

	"gitea.ysicing.net/cloud/pangu/pkg/util"
	"github.com/cockroachdb/errors"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var RDB *redis.Client

func SetCache() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     util.GetKeyFromYaml("cache.host", "127.0.0.1:6379"),
		DB:       util.GetKeyIntFromYaml("cache.db", 6),
		Password: util.GetKeyFromYaml("cache.password", ""),
	})
	_, err := RDB.Ping(context.TODO()).Result()
	if err != nil {
		return errors.Newf("redis cache init failed: %v", err)
	}
	logrus.Info("redis cache init success")
	return nil
}
