package service

import (
	"backend/internal/config"
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/internal/types"
)

type RefreshTokensRequest = types.RefreshTokensRequest

type Users interface {
	CreateUser(user domain.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	VerifyUser(identifier, password, verifyBy string) (domain.User, error)
	GetUsernameByID(userID int64) (string, error)
}

type RefreshSession interface {
	CreateRefreshSession(domain.User, domain.RefreshSession, *config.Config) (string, string, error)
	GetRefreshSession(accessToken string) (domain.RefreshSession, error)
	DeleteRefreshSession(refreshToken string) error
	UpdateRefreshSession(token domain.RefreshToken, req RefreshTokensRequest, cfg *config.Config) (domain.AccessToken, domain.RefreshToken, error)
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
	SendMessage(message domain.Message) error
}

type Service struct {
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

func NewService(repos *repository.Repositories) *Service {
	return &Service{
		Users:          NewUsersService(repos.Users),
		RefreshSession: NewRefreshSessionService(repos.RefreshSession),
		WebSocket:      NewWebSocketService(repos.WebSocket),
		Categories:     NewCategoriesService(repos.Categories),
		Brands:         NewBrandsService(repos.Brands),
		Products:       NewProductsService(repos.Products),
		ProductsImages: NewProductsImagesService(repos.ProductsImages),
		Orders:         NewOrdersService(repos.Orders),
		OrdersContent:  NewOrdersContentService(repos.OrdersContent),
	}
}
