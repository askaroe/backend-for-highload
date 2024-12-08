package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		ProductID int64  `json:"productId"`
		UserID    int64  `json:"userId"`
		Rating    int    `json:"rating"`
		Comment   string `json:"comment"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate rating value
	if reqBody.Rating < 1 || reqBody.Rating > 5 {
		c.JSON(400, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	// Create a new review
	review := models.Review{
		ProductID: reqBody.ProductID,
		UserID:    reqBody.UserID,
		Rating:    reqBody.Rating,
		Comment:   reqBody.Comment,
	}

	result := gorm.DB.Create(&review)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create review"})
		return
	}

	c.JSON(201, gin.H{
		"review": review,
	})
}

func GetAllReviews(c *gin.Context) {
	var reviews []models.Review
	result := gorm.DB.Preload("Product").Preload("User").Find(&reviews)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(200, gin.H{
		"reviews": reviews,
	})
}

func GetReviewById(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	result := gorm.DB.Preload("Product").Preload("User").First(&review, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Review not found"})
		return
	}

	c.JSON(200, gin.H{
		"review": review,
	})
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		ProductID int64  `json:"productId"`
		UserID    int64  `json:"userId"`
		Rating    int    `json:"rating"`
		Comment   string `json:"comment"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Validate rating value
	if reqBody.Rating < 1 || reqBody.Rating > 5 {
		c.JSON(400, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	var review models.Review
	result := gorm.DB.First(&review, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Review not found"})
		return
	}

	// Update review details
	review.ProductID = reqBody.ProductID
	review.UserID = reqBody.UserID
	review.Rating = reqBody.Rating
	review.Comment = reqBody.Comment

	gorm.DB.Save(&review)

	c.JSON(200, gin.H{
		"review": review,
	})
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")

	var review models.Review
	result := gorm.DB.First(&review, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Review not found"})
		return
	}

	// Delete review
	gorm.DB.Delete(&review)

	c.Status(200)
}
