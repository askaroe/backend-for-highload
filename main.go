package main

import (
	"context"
	"github.com/alinadsm04/backend-for-highload/internal/handlers"
	"github.com/alinadsm04/backend-for-highload/internal/middleware"
	"github.com/alinadsm04/backend-for-highload/internal/storage/cache"
	"github.com/alinadsm04/backend-for-highload/internal/storage/gorm"
	"github.com/alinadsm04/backend-for-highload/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"os"
)

var logger log.Logger

func main() {
	ctx := context.Background()

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("connecting to db...")
	gorm.ConnectToDb()
	logger.Println("connected to db")

	logger.Println("connecting to Redis")
	cache.ConnectToRedis(ctx)
	logger.Println("connected to Redis")

	logger.Println("initializing prometheus")
	utils.InitPrometheus()
	logger.Println("initializing prometheus")

	r := gin.Default()

	r.Use(middleware.MetricsMiddleware())

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.POST("/api/auth/login", handlers.Login)
	r.POST("/api/auth/register", handlers.Register)

	api := r.Group("api/v1")
	api.Use(middleware.JWTMiddleware())
	{
		// Users Routes
		api.GET("/users", handlers.GetAllUsers)
		api.GET("/users/:user_id", handlers.GetUserById)
		api.DELETE("/users/:user_id", handlers.DeleteUser)

		// Products Routes
		api.GET("/products", handlers.GetAllProducts)
		api.GET("/products/:product_id", handlers.GetProductById)
		api.POST("/products", handlers.CreateProduct)
		api.PUT("/products/:product_id", handlers.UpdateProduct)
		api.DELETE("/products/:product_id", handlers.DeleteProduct)

		// Categories Routes
		api.GET("/categories", handlers.GetAllCategories)
		api.GET("/categories/:category_id", handlers.GetCategoryById)
		api.POST("/categories", handlers.CreateCategory)
		api.PUT("/categories/:category_id", handlers.UpdateCategory)
		api.DELETE("/categories/:category_id", handlers.DeleteCategory)

		// Orders Routes
		api.GET("/orders", handlers.GetAllOrders)
		api.GET("/orders/:order_id", handlers.GetOrderById)
		api.POST("/orders", handlers.CreateOrder)
		api.PUT("/orders/:order_id", handlers.UpdateOrder)
		api.DELETE("/orders/:order_id", handlers.DeleteOrder)

		// OrderItems Routes
		api.GET("/orderitems", handlers.GetAllOrderItems)
		api.GET("/orderitems/:orderitem_id", handlers.GetOrderItemById)
		api.POST("/orderitems", handlers.CreateOrderItem)
		api.PUT("/orderitems/:orderitem_id", handlers.UpdateOrderItem)
		api.DELETE("/orderitems/:orderitem_id", handlers.DeleteOrderItem)

		// ShoppingCarts Routes
		api.GET("/shoppingcarts", handlers.GetAllShoppingCarts)
		api.GET("/shoppingcarts/:cart_id", handlers.GetShoppingCartById)
		api.POST("/shoppingcarts", handlers.CreateShoppingCart)
		api.PUT("/shoppingcarts/:cart_id", handlers.UpdateShoppingCart)
		api.DELETE("/shoppingcarts/:cart_id", handlers.DeleteShoppingCart)

		// CartItems Routes
		api.GET("/cartitems", handlers.GetAllCartItems)
		api.GET("/cartitems/:cartitem_id", handlers.GetCartItemById)
		api.POST("/cartitems", handlers.CreateCartItem)
		api.PUT("/cartitems/:cartitem_id", handlers.UpdateCartItem)
		api.DELETE("/cartitems/:cartitem_id", handlers.DeleteCartItem)

		// Payments Routes
		api.GET("/payments", handlers.GetAllPayments)
		api.GET("/payments/:payment_id", handlers.GetPaymentById)
		api.POST("/payments", handlers.CreatePayment)
		api.PUT("/payments/:payment_id", handlers.UpdatePayment)
		api.DELETE("/payments/:payment_id", handlers.DeletePayment)

		// Reviews Routes
		api.GET("/reviews", handlers.GetAllReviews)
		api.GET("/reviews/:review_id", handlers.GetReviewById)
		api.POST("/reviews", handlers.CreateReview)
		api.PUT("/reviews/:review_id", handlers.UpdateReview)
		api.DELETE("/reviews/:review_id", handlers.DeleteReview)

		// Wishlists Routes
		api.GET("/wishlists", handlers.GetAllWishlists)
		api.GET("/wishlists/:wishlist_id", handlers.GetWishlistById)
		api.POST("/wishlists", handlers.CreateWishlist)
		api.PUT("/wishlists/:wishlist_id", handlers.UpdateWishlist)
		api.DELETE("/wishlists/:wishlist_id", handlers.DeleteWishlist)

		// WishListItems Routes
		api.GET("/wishlistitems", handlers.GetAllWishListItems)
		api.GET("/wishlistitems/:wishlistitem_id", handlers.GetWishListItemById)
		api.POST("/wishlistitems", handlers.CreateWishListItem)
		api.PUT("/wishlistitems/:wishlistitem_id", handlers.UpdateWishListItem)
		api.DELETE("/wishlistitems/:wishlistitem_id", handlers.DeleteWishListItem)
	}

	err := r.Run()
	if err != nil {
		log.Fatal("error starting server: " + err.Error())
		return
	}

}
