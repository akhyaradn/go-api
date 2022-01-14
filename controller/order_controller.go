package controller

import (
	"api/models"
	"api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrdersIndex(c *gin.Context) {
	orders, err := service.GetAllOrders()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func Post(c *gin.Context) {
	order := models.Order{}
	err := c.ShouldBindJSON(&order)
	ok := order.ValidateCreate()

	if err != nil || !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Unprocessable Entity",
			"message": "invalid json data",
		})
		return
	}

	err = service.InsertOrderAndItems(&order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func OrderIndex(c *gin.Context) {
	orderId, _ := strconv.ParseUint(c.Param("orderId"), 10, 64)

	order, err := service.GetOrderById(orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func Put(c *gin.Context) {
	orderId, _ := strconv.ParseUint(c.Param("orderId"), 10, 64)
	order := models.Order{}
	err := c.ShouldBindJSON(&order)
	ok := order.ValidateUpdate()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal server error",
			"message": err.Error(),
		})
		return
	}

	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "Unprocessable entity",
			"message": "Invalid request body",
		})
		return
	}

	order.Id = uint(orderId)

	err = service.UpdateOrderAndItems(&order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func Delete(c *gin.Context) {
	orderId, _ := strconv.ParseUint(c.Param("orderId"), 10, 64)
	err := service.DeleteOrderAndItems(orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, map[string]interface{}{})
}
