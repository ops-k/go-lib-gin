package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ops-k/go-lib-gin/controllers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

// HealthRoutes struct
type ActuatorRoutes struct {
	logger             zerolog.Logger
	actuatorController controllers.ActuatorController
}

// Setup user routes
func (s ActuatorRoutes) BindTo(engine *gin.Engine) {
	promHandler := promhttp.Handler()

	engine.GET("/health", s.actuatorController.GetHealth)
	engine.GET("/metrics", gin.WrapH(promHandler))

	group := engine.Group("/actuator")
	{
		group.GET("/health", s.actuatorController.GetHealth)
		group.GET("/health/liveness", s.actuatorController.GetHealthLiveness)
		group.GET("/health/readiness", s.actuatorController.GetHealthReadiness)
		group.GET("/info", s.actuatorController.GetInfo)
		group.GET("/metrics", gin.WrapH(promHandler))
	}
}

// NewHealthRoutes creates new user controller
func NewActuatorRoutes(
	logger zerolog.Logger,
	actuatorController controllers.ActuatorController,
) ActuatorRoutes {
	return ActuatorRoutes{
		logger:             logger,
		actuatorController: actuatorController,
	}
}
