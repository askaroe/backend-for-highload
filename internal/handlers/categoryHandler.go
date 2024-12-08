package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		Name     string `json:"name"`
		ParentID int64  `json:"parentId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new category
	category := models.Category{
		Name:     reqBody.Name,
		ParentID: reqBody.ParentID,
	}

	result := gorm.DB.Create(&category)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(201, gin.H{
		"category": category,
	})
}

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	result := gorm.DB.Preload("Parent").Find(&categories)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(200, gin.H{
		"categories": categories,
	})
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	result := gorm.DB.Preload("Parent").First(&category, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(200, gin.H{
		"category": category,
	})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		Name     string `json:"name"`
		ParentID int64  `json:"parentId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var category models.Category
	result := gorm.DB.First(&category, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	// Update category
	category.Name = reqBody.Name
	category.ParentID = reqBody.ParentID

	gorm.DB.Save(&category)

	c.JSON(200, gin.H{
		"category": category,
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	result := gorm.DB.First(&category, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	// Delete category
	gorm.DB.Delete(&category)

	c.Status(200)
}
