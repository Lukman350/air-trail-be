package main

import (
	"air-trail-backend/routers"
	"air-trail-backend/utils/env"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	mode := env.GetEnv("APP_ENV", "debug")
	port := env.GetEnv("APP_PORT", "8080")

	gin.SetMode(mode)
	router := gin.Default()

	cors := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	router.Use(cors)

	aircraftIcons := router.Group("/aircraft")
	aircraftIcons.Use(cors)
	aircraftIcons.Static("/", "./static/icons/aircraft")

	routers.InitRouters(router)

	log.Printf("🚀 Starting application in %s mode on localhost:%s", gin.Mode(), port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
