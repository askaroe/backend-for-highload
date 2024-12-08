package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateWishListItem(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		WishlistID int64 `json:"wishlistId"`
		ProductID  int64 `json:"productId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new wish list item
	wishListItem := models.WishListItem{
		WishlistID: reqBody.WishlistID,
		ProductID:  reqBody.ProductID,
	}

	result := gorm.DB.Create(&wishListItem)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create wish list item"})
		return
	}

	c.JSON(201, gin.H{
		"wishListItem": wishListItem,
	})
}

func GetAllWishListItems(c *gin.Context) {
	var wishListItems []models.WishListItem
	result := gorm.DB.Preload("Wishlist").Preload("Product").Find(&wishListItems)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch wish list items"})
		return
	}

	c.JSON(200, gin.H{
		"wishListItems": wishListItems,
	})
}

func GetWishListItemById(c *gin.Context) {
	id := c.Param("id")
	var wishListItem models.WishListItem
	result := gorm.DB.Preload("Wishlist").Preload("Product").First(&wishListItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Wish list item not found"})
		return
	}

	c.JSON(200, gin.H{
		"wishListItem": wishListItem,
	})
}

func UpdateWishListItem(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		WishlistID int64 `json:"wishlistId"`
		ProductID  int64 `json:"productId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var wishListItem models.WishListItem
	result := gorm.DB.First(&wishListItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Wish list item not found"})
		return
	}

	// Update wish list item details
	wishListItem.WishlistID = reqBody.WishlistID
	wishListItem.ProductID = reqBody.ProductID

	gorm.DB.Save(&wishListItem)

	c.JSON(200, gin.H{
		"wishListItem": wishListItem,
	})
}

func DeleteWishListItem(c *gin.Context) {
	id := c.Param("id")

	var wishListItem models.WishListItem
	result := gorm.DB.First(&wishListItem, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Wish list item not found"})
		return
	}

	// Delete wish list item
	gorm.DB.Delete(&wishListItem)

	c.Status(200)
}
