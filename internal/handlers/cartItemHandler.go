package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateCartItem(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		CartID    int64 `json:"cartId"`
		ProductID int64 `json:"productId"`
		Quantity  int   `json:"quantity"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new cart item
	cartItem := models.CartItem{
		CartID:    reqBody.CartID,
		ProductID: reqBody.ProductID,
		Quantity:  reqBody.Quantity,
	}

	result := gorm.DB.Create(&cartItem)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create cart item"})
		return
	}

	c.JSON(201, gin.H{
		"cartItem": cartItem,
	})
}

func GetAllCartItems(c *gin.Context) {
	var cartItems []models.CartItem
	result := gorm.DB.Preload("Cart").Preload("Product").Find(&cartItems)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	c.JSON(200, gin.H{
		"cartItems": cartItems,
	})
}

func GetCartItemById(c *gin.Context) {
	id := c.Param("id")
	var cartItem models.CartItem
	result := gorm.DB.Preload("Cart").Preload("Product").First(&cartItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Cart item not found"})
		return
	}

	c.JSON(200, gin.H{
		"cartItem": cartItem,
	})
}

func UpdateCartItem(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		CartID    int64 `json:"cartId"`
		ProductID int64 `json:"productId"`
		Quantity  int   `json:"quantity"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var cartItem models.CartItem
	result := gorm.DB.First(&cartItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Cart item not found"})
		return
	}

	// Update cart item
	cartItem.CartID = reqBody.CartID
	cartItem.ProductID = reqBody.ProductID
	cartItem.Quantity = reqBody.Quantity

	gorm.DB.Save(&cartItem)

	c.JSON(200, gin.H{
		"cartItem": cartItem,
	})
}

func DeleteCartItem(c *gin.Context) {
	id := c.Param("id")

	var cartItem models.CartItem
	result := gorm.DB.First(&cartItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Cart item not found"})
		return
	}

	// Delete cart item
	gorm.DB.Delete(&cartItem)

	c.Status(200)
}
