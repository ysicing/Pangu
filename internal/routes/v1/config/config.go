package config

import (
	"gitea.ysicing.net/cloud/pangu/internal/routes"
	"github.com/gin-gonic/gin"
)

type Config struct{}

func init() {
	routes.Migrate(NewHandler)
}

func NewHandler() (routes.RouteRegister, error) {
	return &Config{}, nil
}

func (api Config) ApplyRoute(r *gin.Engine) {
	c := r.Group("/api/configs")
	c.GET("", api.Get)
}

func (api Config) Name() string {
	return "config"
}

// Get
// @Summary 获取配置
// @Tags config
// @Accept application/json
// @Param Authorization header string false "jwtToken"
// @Param X-Auth-Token header string false "staticToken"
// @Security ApiKeyAuth
// @Success 200
// @Router /api/configs [get]
func (api Config) Get(c *gin.Context) {
}
