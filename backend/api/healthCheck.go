package api

import (
	"marketeer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	var response models.GenericResponse

	response.Code = http.StatusOK
	response.Status = "success"
	//response.Message = "running in " + config.GetEnvConfig("ENVIRONMENT") + " server..."
	response.Message = "START DATE: November 24,2022 (1:40 PM)"
	c.JSON(http.StatusOK, response)
}
