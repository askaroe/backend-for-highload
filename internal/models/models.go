package models

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updated_at"`
	Username  string    `gorm:"size:50;not null" json:"username"`
	Email     string    `gorm:"size:50;not null" json:"email"`
	Password  string    `gorm:"size:50;not null" json:"password"`
	FirstName string    `gorm:"size:50" json:"firstName"`
	LastName  string    `gorm:"size:50" json:"lastName"`
}

func (User) TableName() string {
	return "users"
}

type Product struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:100;not null" json:"name"`
	Description   string    `gorm:"type:text" json:"description"`
	Price         float64   `gorm:"type:numeric(10,2);not null" json:"price"`
	StockQuantity int       `gorm:"not null;default:0" json:"stockQuantity"`
	CategoryID    int64     `gorm:"not null" json:"categoryId"`
	CreatedAt     time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	Category      Category       `gorm:"foreignKey:CategoryID" json:"category"`
	OrderItems    []OrderItem    `gorm:"foreignKey:ProductID" json:"orderItems"`
	Reviews       []Review       `gorm:"foreignKey:ProductID" json:"reviews"`
	WishListItems []WishListItem `gorm:"foreignKey:ProductID" json:"wishListItems"`
	CartItems     []CartItem     `gorm:"foreignKey:ProductID" json:"cartItems"`
}

func (Product) TableName() string {
	return "product"
}

type Category struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	ParentID  int64     `gorm:"default:null" json:"parentId"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`
	Parent    *Category `gorm:"foreignKey:ParentID" json:"parent"`
	Products  []Product `gorm:"foreignKey:CategoryID" json:"products"`
}

func (Category) TableName() string {
	return "category"
}

type Order struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int64     `gorm:"not null" json:"userId"`
	OrderStatus string    `gorm:"size:50;not null" json:"orderStatus"`
	TotalAmount float64   `gorm:"type:numeric(10,2);not null" json:"totalAmount"`
	CreatedAt   time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	User       User        `gorm:"foreignKey:UserID" json:"user"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"orderItems"`
	Payments   []Payment   `gorm:"foreignKey:OrderID" json:"payments"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int64     `gorm:"not null" json:"orderId"`
	ProductID int64     `gorm:"not null" json:"productId"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"type:numeric(10,2);not null" json:"price"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	Order   Order   `gorm:"foreignKey:OrderID" json:"order"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}

func (OrderItem) TableName() string {
	return "orderitems"
}

type ShoppingCart struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"not null" json:"userId"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	User      User       `gorm:"foreignKey:UserID" json:"user"`
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cartItems"`
}

func (ShoppingCart) TableName() string {
	return "shoppingcart"
}

type CartItem struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CartID    int64     `gorm:"not null" json:"cartId"`
	ProductID int64     `gorm:"not null" json:"productId"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	Cart    ShoppingCart `gorm:"foreignKey:CartID" json:"cart"`
	Product Product      `gorm:"foreignKey:ProductID" json:"product"`
}

func (CartItem) TableName() string {
	return "cartitems"
}

type Payment struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID       int64     `gorm:"not null" json:"orderId"`
	PaymentMethod string    `gorm:"size:50;not null" json:"paymentMethod"`
	Amount        float64   `gorm:"type:numeric(10,2);not null" json:"amount"`
	Status        string    `gorm:"size:50;not null" json:"status"`
	CreatedAt     time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	Order Order `gorm:"foreignKey:OrderID" json:"order"`
}

func (Payment) TableName() string {
	return "payments"
}

type Review struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int64     `gorm:"not null" json:"productId"`
	UserID    int64     `gorm:"not null" json:"userId"`
	Rating    int       `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`
	User    User    `gorm:"foreignKey:UserID" json:"user"`
}

func (Review) TableName() string {
	return "reviews"
}

type Wishlist struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"not null" json:"userId"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:now()" json:"updatedAt"`

	User          User           `gorm:"foreignKey:UserID" json:"user"`
	WishListItems []WishListItem `gorm:"foreignKey:WishlistID" json:"wishListItems"`
}

func (Wishlist) TableName() string {
	return "wichlists"
}

type WishListItem struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	WishlistID int64     `gorm:"not null" json:"wishlistId"`
	ProductID  int64     `gorm:"not null" json:"productId"`
	CreatedAt  time.Time `gorm:"type:timestamp with time zone;default:now()" json:"createdAt"`

	Wishlist Wishlist `gorm:"foreignKey:WishlistID" json:"wishlist"`
	Product  Product  `gorm:"foreignKey:ProductID" json:"product"`
}

func (WishListItem) TableName() string {
	return "wishlistitem"
}
