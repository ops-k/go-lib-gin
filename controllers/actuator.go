package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	actuator_models "github.com/ops-k/go-lib-actuator/models"
	actuator_services "github.com/ops-k/go-lib-actuator/services"
	"github.com/rs/zerolog"
)

// ActuatorController struct
type ActuatorController struct {
	logger                zerolog.Logger
	actuatorHealthService *actuator_services.ActuatorHealthService
	actuatorInfoService   *actuator_services.ActuatorInfoService
}

// Responds to a healthcheck
func (ctrl ActuatorController) GetHealth(c *gin.Context) {
	response := ctrl.actuatorHealthService.GetHealth(c.Request.Context())
	if response.Status == actuator_models.HealthStatusDown || response.Status == actuator_models.HealthStatusOutOfService {
		c.JSON(http.StatusServiceUnavailable, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// Responds to a healthcheck
func (ctrl ActuatorController) GetHealthLiveness(c *gin.Context) {
	response := ctrl.actuatorHealthService.GetHealthLiveness(c.Request.Context())
	if response.Status == actuator_models.HealthStatusDown || response.Status == actuator_models.HealthStatusOutOfService {
		c.JSON(http.StatusServiceUnavailable, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// Responds to a healthcheck
func (ctrl ActuatorController) GetHealthReadiness(c *gin.Context) {
	response := ctrl.actuatorHealthService.GetHealthReadiness(c.Request.Context())
	if response.Status == actuator_models.HealthStatusDown || response.Status == actuator_models.HealthStatusOutOfService {
		c.JSON(http.StatusServiceUnavailable, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

// returns information about the deployment
func (ctrl ActuatorController) GetInfo(c *gin.Context) {
	response := ctrl.actuatorInfoService.GetInfo(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

// Creates the ActuatorController
func NewActuatorController(
	logger zerolog.Logger,
	actuatorHealthService *actuator_services.ActuatorHealthService,
	actuatorInfoService *actuator_services.ActuatorInfoService,
) ActuatorController {
	return ActuatorController{
		logger:                logger,
		actuatorHealthService: actuatorHealthService,
		actuatorInfoService:   actuatorInfoService,
	}
}
