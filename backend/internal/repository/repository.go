package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	GetUserByUsername(username, password string) (domain.User, error)
	GetUserByEmail(username, password string) (domain.User, error)
	CreateUser(user domain.User) (int64, error)
	GetUsernameByID(userID int64) (string, error)
}

type RefreshSession interface {
	CreateRefreshSession(session domain.RefreshSession) (domain.RefreshSession, error)
	DeleteRefreshSession(token domain.RefreshToken) (domain.RefreshSession, error)
}

type Categories interface {
	GetIDByCategory(name string) (int64, error)
	GetCategoryByID(categoryID int64) (string, error)
	CreateCategory(name string) (int64, error)
	DeleteCategory(categoryID int64) error
	UpdateCategory(categoryID int64, name string) error
}

type Brands interface {
	GetIDByBrand(name string) (int64, error)
	GetBrandByID(categoryID int64) (string, error)
	CreateBrand(name string) (int64, error)
	DeleteBrand(categoryID int64) error
	UpdateBrand(categoryID int64, name string) error
}

type Products interface {
	CreateProduct(product domain.Product) (int64, error)
	Get(name string, brandID int64, categoryID int64) (domain.Product, error)
	GetAllByCategory(categoryID int64) ([]domain.Product, error)
	GetAllByName(name string) ([]domain.Product, error)
	GetAllByBrand(brandID int64) ([]domain.Product, error)
	UpdateProduct(product domain.Product) error
	DeleteProduct(productID int64) error
}

type ProductsImages interface {
	GetImageHashByProductID(productID int64) (string, error)
	CreateProductImage(productImage domain.ProductImage) (int64, error)
	UpdateProductImage(oldName string, productImage domain.ProductImage) error
	DeleteProductImageByName(name string) error
	DeleteProductImageByID(imageID int64) error
}

type Orders interface {
	CreateOrder(order domain.Order) (int64, error)
	GetOrderByID(orderID int64) (domain.Order, error)
	GetOrdersByUserID(userID int64) ([]domain.Order, error)
	UpdateOrder(order domain.Order) error
	DeleteOrder(orderID int64) error
}

type OrdersContent interface {
	CreateOrderContent(orderContent domain.OrderContent) (int64, error)
	UpdateOrderContent(orderContent domain.OrderContent) error
	DeleteOrderContent(orderContentID int64) error
}

type WebSocket interface {
	CreateMessage(message domain.Message) error
}

type Repositories struct {
	Users          Users
	RefreshSession RefreshSession
	WebSocket      WebSocket
	Categories     Categories
	Brands         Brands
	Products       Products
	ProductsImages ProductsImages
	Orders         Orders
	OrdersContent  OrdersContent
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:          NewUsersRepo(db),
		RefreshSession: NewRefreshSessionRepo(db),
		WebSocket:      NewWebSocketRepo(db),
		Categories:     NewCategoriesRepo(db),
		Brands:         NewBrandsRepo(db),
		Products:       NewProductsRepo(db),
		ProductsImages: NewProductsImagesRepo(db),
		Orders:         NewOrdersRepo(db),
		OrdersContent:  NewOrdersContentRepo(db),
	}
}
