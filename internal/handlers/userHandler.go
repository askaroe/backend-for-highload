package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := gorm.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := gorm.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		Username  string `json:"username"`
		Firstname string `json:"firstName"`
		Lastname  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var user models.User
	result := gorm.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Update user
	user.Username = reqBody.Username
	user.FirstName = reqBody.Firstname
	user.LastName = reqBody.Lastname
	user.Email = reqBody.Email
	user.Password = reqBody.Password

	gorm.DB.Save(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := gorm.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Delete user
	gorm.DB.Delete(&user)

	c.Status(200)
}
