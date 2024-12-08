package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateOrderItem(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		OrderID   int64   `json:"orderId"`
		ProductID int64   `json:"productId"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new order item
	orderItem := models.OrderItem{
		OrderID:   reqBody.OrderID,
		ProductID: reqBody.ProductID,
		Quantity:  reqBody.Quantity,
		Price:     reqBody.Price,
	}

	result := gorm.DB.Create(&orderItem)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create order item"})
		return
	}

	c.JSON(201, gin.H{
		"orderItem": orderItem,
	})
}

func GetAllOrderItems(c *gin.Context) {
	var orderItems []models.OrderItem
	result := gorm.DB.Preload("Order").Preload("Product").Find(&orderItems)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch order items"})
		return
	}

	c.JSON(200, gin.H{
		"orderItems": orderItems,
	})
}

func GetOrderItemById(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem
	result := gorm.DB.Preload("Order").Preload("Product").First(&orderItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order item not found"})
		return
	}

	c.JSON(200, gin.H{
		"orderItem": orderItem,
	})
}

func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		OrderID   int64   `json:"orderId"`
		ProductID int64   `json:"productId"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var orderItem models.OrderItem
	result := gorm.DB.First(&orderItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order item not found"})
		return
	}

	// Update order item
	orderItem.OrderID = reqBody.OrderID
	orderItem.ProductID = reqBody.ProductID
	orderItem.Quantity = reqBody.Quantity
	orderItem.Price = reqBody.Price

	gorm.DB.Save(&orderItem)

	c.JSON(200, gin.H{
		"orderItem": orderItem,
	})
}

func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")

	var orderItem models.OrderItem
	result := gorm.DB.First(&orderItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order item not found"})
		return
	}

	// Delete order item
	gorm.DB.Delete(&orderItem)

	c.Status(200)
}
