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
