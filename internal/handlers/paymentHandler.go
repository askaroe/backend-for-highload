package handlers

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alinadsm04/backend-for-highload/internal/models"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreatePayment(c *gin.Context) {
	// Getting data from request body
	var reqBody struct {
		OrderID       int64   `json:"orderId"`
		PaymentMethod string  `json:"paymentMethod"`
		Amount        float64 `json:"amount"`
		Status        string  `json:"status"`
	}

	if err := c.Bind(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a new payment
	payment := models.Payment{
		OrderID:       reqBody.OrderID,
		PaymentMethod: reqBody.PaymentMethod,
		Amount:        reqBody.Amount,
		Status:        reqBody.Status,
	}

	brokers := []string{"broker01:9092"}
	topic := "payments"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Initialize producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start Kafka producer"})
		return
	}
	defer producer.Close()

	data, err := json.Marshal(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling payment data"})
		return
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to Kafka " + err.Error()})
		return
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	c.JSON(200, gin.H{"message": "Payment created and sent to Kafka, check consumer"})
}

func GetAllPayments(c *gin.Context) {
	var payments []models.Payment
	result := gorm.DB.Preload("Order").Find(&payments)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch payments"})
		return
	}

	c.JSON(200, gin.H{
		"payments": payments,
	})
}

func GetPaymentById(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	result := gorm.DB.Preload("Order").First(&payment, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(200, gin.H{
		"payment": payment,
	})
}

func UpdatePayment(c *gin.Context) {
	id := c.Param("id")

	var reqBody struct {
		OrderID       int64   `json:"orderId"`
		PaymentMethod string  `json:"paymentMethod"`
		Amount        float64 `json:"amount"`
		Status        string  `json:"status"`
	}

	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	var payment models.Payment
	result := gorm.DB.First(&payment, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Payment not found"})
		return
	}

	// Update payment details
	payment.OrderID = reqBody.OrderID
	payment.PaymentMethod = reqBody.PaymentMethod
	payment.Amount = reqBody.Amount
	payment.Status = reqBody.Status

	gorm.DB.Save(&payment)

	c.JSON(200, gin.H{
		"payment": payment,
	})
}

func DeletePayment(c *gin.Context) {
	id := c.Param("id")

	var payment models.Payment
	result := gorm.DB.First(&payment, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Payment not found"})
		return
	}

	// Delete payment
	gorm.DB.Delete(&payment)

	c.Status(200)
}
