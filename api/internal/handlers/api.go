package handlers

import (
	"net/http"

	"github.com/Vladisluv228/ApiBdGoExample/api/internal/models"
	"github.com/gin-gonic/gin"
)

// GET /logs
func GetLogsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	Logs, err := models.GetLogs(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Logs)
}

// POST /logs
func CreateLogHandler(c *gin.Context) {
	var input models.LogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	log, err := models.CreateLog(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}