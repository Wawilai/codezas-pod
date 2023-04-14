package main

import (
	"codezas-pos/db"
	"codezas-pos/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
	}

	db.ConnectDB()
	db.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	os.MkdirAll("uploads/products", 0755)
	r := gin.Default()
	r.Use(cors.New(corsConfig))
	r.Static("/uploads", "./uploads")
	routes.ServRoutes(r)
	r.Run(":" + os.Getenv("PORT"))
}
