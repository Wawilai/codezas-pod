package routes

import (
	"codezas-pos/controllers"

	"github.com/gin-gonic/gin"
)

func ServRoutes(route *gin.Engine) {

	productController := controllers.Product{}
	rpg := route.Group("/products")
	rpg.GET("", productController.FindAll)
	rpg.GET("/:id", productController.FindOne)
	rpg.POST("", productController.Create)
	rpg.PATCH("/:id", productController.Update)
	rpg.DELETE("/:id", productController.Delete)

	categoryController := controllers.Category{}
	cpg := route.Group("/categories")
	cpg.GET("", categoryController.FindAll)
	cpg.GET("/:id", categoryController.FindOne)
	cpg.POST("", categoryController.Create)
	cpg.PATCH("/:id", categoryController.Update)
	cpg.DELETE("/:id", categoryController.Delete)

	orderController := controllers.Order{}
	odg := route.Group("/orders")
	odg.GET("", orderController.FindAll)
	odg.GET("/:id", orderController.FindOne)
	odg.POST("", orderController.Create)
}
