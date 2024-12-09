package handlers

import (
	"encoding/json"
	"errors"
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/cache"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

func CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var reqBody struct {
		Name          string  `json:"name"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		StockQuantity int     `json:"stockQuantity"`
		CategoryID    int64   `json:"categoryId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new product
	product := models.Product{
		Name:          reqBody.Name,
		Description:   reqBody.Description,
		Price:         reqBody.Price,
		StockQuantity: reqBody.StockQuantity,
		CategoryID:    reqBody.CategoryID,
	}

	result := gorm.DB.Create(&product)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}

	// Clear product cache after creating a new product (optional)
	cache.Rdb.Del(ctx, "products_list")

	c.JSON(201, gin.H{
		"product": product,
	})
}

// GetAllProducts fetches all products and uses Redis cache
func GetAllProducts(c *gin.Context) {
	ctx := c.Request.Context()

	// Check if products are cached
	cachedProducts, err := cache.Rdb.Get(ctx, "products_list").Result()
	if errors.Is(err, redis.Nil) {
		// Cache miss, fetch from DB
		var products []models.Product
		result := gorm.DB.Preload("Category").Find(&products)

		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch products"})
			return
		}

		// Cache the result for 10 minutes
		cachedData, _ := json.Marshal(products)
		cache.Rdb.Set(ctx, "products_list", cachedData, 10*time.Minute)

		// Send response
		c.JSON(200, gin.H{
			"products": products,
		})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching from Redis"})
		return
	}

	// Cache hit, return cached data
	var products []models.Product
	if err := json.Unmarshal([]byte(cachedProducts), &products); err != nil {
		c.JSON(500, gin.H{"error": "Failed to unmarshal cached data"})
		return
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}

// GetProductById fetches a product by ID and uses Redis cache
func GetProductById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	cacheKey := "product_" + id

	// Check if product is cached
	cachedProduct, err := cache.Rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss, fetch from DB
		var product models.Product
		result := gorm.DB.Preload("Category").First(&product, id)

		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}

		// Cache the result for 10 minutes
		cachedData, _ := json.Marshal(product)
		cache.Rdb.Set(ctx, cacheKey, cachedData, 10*time.Minute)

		// Send response
		c.JSON(200, gin.H{
			"product": product,
		})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching from Redis"})
		return
	}

	// Cache hit, return cached data
	var product models.Product
	if err := json.Unmarshal([]byte(cachedProduct), &product); err != nil {
		c.JSON(500, gin.H{"error": "Failed to unmarshal cached data"})
		return
	}

	c.JSON(200, gin.H{
		"product": product,
	})
}

// UpdateProduct updates a product (cache is invalidated after update)
func UpdateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var reqBody struct {
		Name          string  `json:"name"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		StockQuantity int     `json:"stockQuantity"`
		CategoryID    int64   `json:"categoryId"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var product models.Product
	result := gorm.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	// Update product
	product.Name = reqBody.Name
	product.Description = reqBody.Description
	product.Price = reqBody.Price
	product.StockQuantity = reqBody.StockQuantity
	product.CategoryID = reqBody.CategoryID

	gorm.DB.Save(&product)

	// Clear the relevant cache after updating product
	cache.Rdb.Del(ctx, "products_list")
	cache.Rdb.Del(ctx, "product_"+id)

	c.JSON(200, gin.H{
		"product": product,
	})
}

// DeleteProduct deletes a product (cache is invalidated after deletion)
func DeleteProduct(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var product models.Product
	result := gorm.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	// Delete product
	gorm.DB.Delete(&product)

	// Clear the relevant cache after deleting product
	cache.Rdb.Del(ctx, "products_list")
	cache.Rdb.Del(ctx, "product_"+id)

	c.Status(200)
}
