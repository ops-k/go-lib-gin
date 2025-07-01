package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	actuator_services "github.com/ops-k/go-lib-actuator/services"
	"github.com/rs/zerolog"
)

// PingController struct
type PingController struct {
	logger      zerolog.Logger
	pingService *actuator_services.PingService
}

// Responds to a ping
func (ctrl PingController) FindAll(c *gin.Context) {
	ctx := c.Request.Context()
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("received ping request.")
	response := ctrl.pingService.Ping(ctx)
	c.JSON(http.StatusOK, response)
}

// Creates the PingController
func NewPingController(
	logger zerolog.Logger,
	pingService *actuator_services.PingService,
) PingController {
	return PingController{
		logger:      logger,
		pingService: pingService,
	}
}
