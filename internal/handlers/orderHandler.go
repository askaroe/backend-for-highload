package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		UserID      int64   `json:"userId"`
		OrderStatus string  `json:"orderStatus"`
		TotalAmount float64 `json:"totalAmount"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new order
	order := models.Order{
		UserID:      reqBody.UserID,
		OrderStatus: reqBody.OrderStatus,
		TotalAmount: reqBody.TotalAmount,
	}

	result := gorm.DB.Create(&order)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(201, gin.H{
		"order": order,
	})
}

func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	result := gorm.DB.Preload("User").Preload("OrderItems").Preload("Payments").Find(&orders)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	result := gorm.DB.Preload("User").Preload("OrderItems").Preload("Payments").First(&order, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, gin.H{
		"order": order,
	})
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		UserID      int64   `json:"userId"`
		OrderStatus string  `json:"orderStatus"`
		TotalAmount float64 `json:"totalAmount"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var order models.Order
	result := gorm.DB.First(&order, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	// Update order
	order.UserID = reqBody.UserID
	order.OrderStatus = reqBody.OrderStatus
	order.TotalAmount = reqBody.TotalAmount

	gorm.DB.Save(&order)

	c.JSON(200, gin.H{
		"order": order,
	})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	var order models.Order
	result := gorm.DB.First(&order, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	// Delete order
	gorm.DB.Delete(&order)

	c.Status(200)
}
