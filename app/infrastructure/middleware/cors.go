package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/k8s-go-sample/config"
)

type Cors struct {
	Config gin.HandlerFunc
}

func NewCors(cfg *config.Config) *Cors {
	c := &Cors{}
	// DefaultConfig returns a generic default configuration mapped to localhost.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Cors.AllowOringins
	corsConfig.AllowHeaders = []string{"authorization", "content-type"}
	corsConfig.ExposeHeaders = []string{"Authorization"}
	corsConfig.AllowCredentials = true

	c.Config = cors.New(corsConfig)

	return c
}
