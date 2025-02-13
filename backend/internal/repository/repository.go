package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	GetUserByUsername(username, password string) (domain.User, error)
	GetUserByEmail(username, password string) (domain.User, error)
	CreateUser(user domain.User) (int64, error)
	GetUsernameById(userId int64) (string, error)
}

type RefreshSession interface {
	CreateRefreshSession(session domain.RefreshSession) (domain.RefreshSession, error)
	DeleteRefreshSession(token domain.RefreshToken) (domain.RefreshSession, error)
}

type Categories interface {
	GetIdByCategory(name string) (int64, error)
	GetCategoryById(categoryId int64) (string, error)
	CreateCategory(name string) (int64, error)
	DeleteCategory(categoryId int64) error
	UpdateCategory(categoryId int64, name string) error
}

type Brands interface {
	GetIdByBrand(name string) (int64, error)
	GetBrandById(brandId int64) (string, error)
	CreateBrand(name string) (int64, error)
	DeleteBrand(brandId int64) error
	UpdateBrand(brandId int64, name string) error
}

type Products interface {
	CreateProduct(product domain.Product) (int64, error)
	Get(name string, brandId int64, categoryId int64) (domain.Product, error)
	GetAllByCategory(categoryId int64) ([]domain.Product, error)
	GetAllByName(name string) ([]domain.Product, error)
	GetAllByBrand(brandId int64) ([]domain.Product, error)
	UpdateProduct(product domain.Product) error
	DeleteProduct(productID int64) error
}

type ProductsImages interface {
	GetSequenceByProductId(productId int64) (int64, error)
	GetImageIdByHash(imageHash string) (int64, error)
	GetImageHashByImageId(imageId int64) (string, error)
	GetImageHashesByProductId(productId int64) ([]string, error)
	CreateProductImage(productImage domain.ProductImage) (int64, error)
	UpdateProductImage(oldName string, newName string) error
	DeleteProductImageByName(name string) error
	DeleteProductImageById(imageId int64) error
}

type Orders interface {
	CreateOrder(order domain.Order) (int64, error)
	GetOrderById(orderId int64) (domain.Order, error)
	GetOrdersByUserId(userId int64) ([]domain.Order, error)
	UpdateOrder(order domain.Order) error
	DeleteOrder(orderId int64) error
}

type OrdersContent interface {
	CreateOrderContent(orderContent domain.OrderContent) (int64, error)
	UpdateOrderContent(orderContent domain.OrderContent) error
	DeleteOrderContent(orderContentId int64) error
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
