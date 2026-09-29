package handlers

import (
	"net/http"

	"celebration-hall/database"
	"celebration-hall/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	var events []models.Event
	database.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}

func GetEventByID(c *gin.Context) {
	var event models.Event
	if err := database.DB.First(&event, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}