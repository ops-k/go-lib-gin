package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ops-k/go-lib-gin/controllers"
	"github.com/rs/zerolog"
)

// PingRoutes struct
type PingRoutes struct {
	logger         zerolog.Logger
	pingController controllers.PingController
}

// Setup user routes
func (s PingRoutes) BindTo(engine *gin.Engine) {
	engine.GET("/ping", s.pingController.FindAll)
}

// NewPingRoutes creates new user controller
func NewPingRoutes(
	logger zerolog.Logger,
	pingController controllers.PingController,
) PingRoutes {
	return PingRoutes{
		logger:         logger,
		pingController: pingController,
	}
}
