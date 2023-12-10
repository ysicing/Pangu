// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RegisterRouteFactory func() (RouteRegister, error)

var Migrates []RegisterRouteFactory

func Migrate(obj RegisterRouteFactory) {
	Migrates = append(Migrates, obj)
}

type RouteRegister interface {
	ApplyRoute(r *gin.Engine)
	Name() string
}

func SetupRoutes(g *gin.Engine) {
	for i := range Migrates {
		h, err := Migrates[i]()
		if err != nil {
			panic(err)
		}
		h.ApplyRoute(g)
		logrus.Debugf("load router %s success...", h.Name())
	}
	logrus.Infof("load %d router success...", len(Migrates))
}
