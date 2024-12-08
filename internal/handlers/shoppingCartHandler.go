package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateShoppingCart(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		UserID int64 `json:"userId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new shopping cart
	shoppingCart := models.ShoppingCart{
		UserID: reqBody.UserID,
	}

	result := gorm.DB.Create(&shoppingCart)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create shopping cart"})
		return
	}

	c.JSON(201, gin.H{
		"shoppingCart": shoppingCart,
	})
}

func GetAllShoppingCarts(c *gin.Context) {
	var shoppingCarts []models.ShoppingCart
	result := gorm.DB.Preload("User").Preload("CartItems").Find(&shoppingCarts)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch shopping carts"})
		return
	}

	c.JSON(200, gin.H{
		"shoppingCarts": shoppingCarts,
	})
}

func GetShoppingCartById(c *gin.Context) {
	id := c.Param("id")
	var shoppingCart models.ShoppingCart
	result := gorm.DB.Preload("User").Preload("CartItems").First(&shoppingCart, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Shopping cart not found"})
		return
	}

	c.JSON(200, gin.H{
		"shoppingCart": shoppingCart,
	})
}

func UpdateShoppingCart(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		UserID int64 `json:"userId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var shoppingCart models.ShoppingCart
	result := gorm.DB.First(&shoppingCart, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Shopping cart not found"})
		return
	}

	// Update shopping cart
	shoppingCart.UserID = reqBody.UserID

	gorm.DB.Save(&shoppingCart)

	c.JSON(200, gin.H{
		"shoppingCart": shoppingCart,
	})
}

func DeleteShoppingCart(c *gin.Context) {
	id := c.Param("id")

	var shoppingCart models.ShoppingCart
	result := gorm.DB.First(&shoppingCart, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Shopping cart not found"})
		return
	}

	// Delete shopping cart
	gorm.DB.Delete(&shoppingCart)

	c.Status(200)
}
