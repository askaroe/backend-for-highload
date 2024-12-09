package handlers

import (
	"github.com/alinadsm04/backend-for-highload/internal/middleware"
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create a new user
	user := models.User{
		Username:  reqBody.Username,
		FirstName: reqBody.Firstname,
		LastName:  reqBody.Lastname,
		Email:     reqBody.Email,
		Password:  string(hashedPassword),
	}

	result := gorm.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var reqBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var dbUser models.User
	if err := gorm.DB.Where("username = ?", reqBody.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(reqBody.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
