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
	GetUsernameById(userId int64) (string, error)
}

type RefreshSession interface {
	CreateRefreshSession(domain.User, domain.RefreshSession, *config.Config) (string, string, error)
	GetRefreshSession(accessToken string) (domain.RefreshSession, error)
	DeleteRefreshSession(refreshToken string) error
	UpdateRefreshSession(token domain.RefreshToken, req RefreshTokensRequest, cfg *config.Config) (domain.AccessToken, domain.RefreshToken, error)
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
	CreateProduct(description string, name string, price float64, categoryId int64, brandId int64) (int64, error)
	Get(name string, brandId int64, categoryId int64) (domain.Product, error)
	GetAllByCategory(categoryId int64) ([]domain.Product, error)
	GetAllByName(name string) ([]domain.Product, error)
	GetAllByBrand(brandId int64) ([]domain.Product, error)
	UpdateProductById(productId int64, description string, name string, price float64, categoryId int64, brandId int64) error
	DeleteProduct(productId int64) error
}

type ProductsImages interface {
	GetSequenceByProductId(productId int64) (int64, error)
	GetImageIdByHash(imageHash string) (int64, error)
	GetImageHashByImageId(imageId int64) (string, error)
	GetImageHashesByProductId(productId int64) ([]string, error)
	CreateProductImage(productId int64, hashString string) (int64, error)
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
	// TODO: GetOrderContent(orderId int64) ([]domain.OrderContent, error)
	CreateOrderContent(orderContent domain.OrderContent) (int64, error)
	UpdateOrderContent(orderContent domain.OrderContent) error
	DeleteOrderContent(orderContentId int64) error
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
