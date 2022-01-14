package router

import (
	"api/controller"
	"api/database"

	"github.com/gin-gonic/gin"
)

func Init() {
	database.StartDB()
}

func StartRouter() *gin.Engine {
	r := gin.Default()

	ordersEndpoint := r.Group("/orders")
	{
		ordersEndpoint.GET("/", controller.OrdersIndex)
		ordersEndpoint.POST("/", controller.Post)
		ordersEndpoint.GET("/:orderId", controller.OrderIndex)
		ordersEndpoint.PUT("/:orderId", controller.Put)
		ordersEndpoint.DELETE("/:orderId", controller.Delete)
	}

	return r
}
