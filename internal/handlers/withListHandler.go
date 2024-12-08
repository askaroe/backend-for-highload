package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateWishlist(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		UserID int64 `json:"userId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new wishlist
	wishlist := models.Wishlist{
		UserID: reqBody.UserID,
	}

	result := gorm.DB.Create(&wishlist)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create wishlist"})
		return
	}

	c.JSON(201, gin.H{
		"wishlist": wishlist,
	})
}

func GetAllWishlists(c *gin.Context) {
	var wishlists []models.Wishlist
	result := gorm.DB.Preload("User").Preload("WishListItems").Find(&wishlists)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch wishlists"})
		return
	}

	c.JSON(200, gin.H{
		"wishlists": wishlists,
	})
}

func GetWishlistById(c *gin.Context) {
	id := c.Param("id")
	var wishlist models.Wishlist
	result := gorm.DB.Preload("User").Preload("WishListItems").First(&wishlist, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Wishlist not found"})
		return
	}

	c.JSON(200, gin.H{
		"wishlist": wishlist,
	})
}

func UpdateWishlist(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		UserID int64 `json:"userId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var wishlist models.Wishlist
	result := gorm.DB.First(&wishlist, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Wishlist not found"})
		return
	}

	// Update wishlist details
	wishlist.UserID = reqBody.UserID

	gorm.DB.Save(&wishlist)

	c.JSON(200, gin.H{
		"wishlist": wishlist,
	})
}

func DeleteWishlist(c *gin.Context) {
	id := c.Param("id")

	var wishlist models.Wishlist
	result := gorm.DB.First(&wishlist, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Wishlist not found"})
		return
	}

	// Delete wishlist
	gorm.DB.Delete(&wishlist)

	c.Status(200)
}
