package main

import (
	"celebration-hall/config"
	"celebration-hall/database"
	"celebration-hall/handlers"
	"celebration-hall/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		api.GET("/events", handlers.GetEvents)
		api.GET("/events/:id", handlers.GetEventByID)

		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			protected.POST("/bookings", handlers.CreateBooking)
			protected.GET("/bookings/my", handlers.GetMyBookings)
		}
	}

	r.Run(":" + config.Port)
}