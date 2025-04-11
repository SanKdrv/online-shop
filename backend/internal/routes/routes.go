package routes

import (
	"backend/internal/config"
	"backend/internal/handlers/websocket"
	"backend/internal/service"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
)

type Handler struct {
	services *service.Service
	hub      *websocket.Hub
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
		hub:      websocket.NewHub(),
	}
}

//func (h *Handler) InitRoutes() *gin.Engine {

// RegisterRoutes настраивает маршруты для приложения.
func (h *Handler) RegisterRoutes(router *chi.Mux, log *slog.Logger, cfg *config.Config) {
	go h.hub.Run()
	// Swagger
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	// API routes
	router.Route("/api/", func(r chi.Router) {
		// Auth
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-in", h.signIn(log, cfg))
			r.Post("/sign-up", h.signUp(log))
			r.Post("/refresh-tokens", h.refreshTokens(log, cfg))
			r.Post("/sign-out", h.signOut(log))
		})

		// User
		r.Route("/user", func(r chi.Router) {
			r.Post("/get-username-by-id", h.getUsernameById(log))
		})

		// Brand
		r.Route("/brand", func(r chi.Router) {
			r.Get("/get-all-brands", h.getAllBrands(log))
			r.Get("/get-id-by-brand", h.getIdByBrand(log))
			r.Get("/get-brand-by-id", h.getBrandById(log))
			r.Post("/create-brand", h.createBrand(log))
			r.Put("/update-brand", h.updateBrand(log))
			r.Delete("/delete-brand", h.deleteBrand(log))
		})

		// Category
		r.Route("/category", func(r chi.Router) {
			r.Get("/get-all-categories", h.getAllCategories(log))
			r.Get("/get-id-by-category", h.getIdByCategory(log))
			r.Get("/get-category-by-id", h.getCategoryById(log))
			r.Post("/create-category", h.createCategory(log))
			r.Delete("/delete-category", h.deleteCategory(log))
			r.Put("/update-category", h.updateCategory(log))
		})

		// Order
		r.Route("/order", func(r chi.Router) {
			r.Get("/get-all-orders", h.getAllOrders(log))
			r.Post("/create-order", h.createOrder(log))
			r.Get("/get-order-by-id", h.getOrderById(log))
			r.Get("/get-orders-by-user-id", h.getOrdersByUserId(log))
			r.Put("/update-order", h.updateOrder(log))
			r.Delete("/delete-order", h.deleteOrder(log))
		})

		// OrderContent
		r.Route("/order-content", func(r chi.Router) {
			r.Post("/create-order-content", h.createOrderContent(log))
			r.Put("/update-order-content", h.updateOrderContent(log))
			r.Delete("/delete-order-content", h.deleteOrderContent(log))
		})

		// Product
		r.Route("/product", func(r chi.Router) {
			r.Post("/create-product", h.createProduct(log))
			r.Get("/get-all-products", h.getAllProducts(log))
			r.Get("/get-product", h.getProduct(log))
			r.Get("/get-all-by-category", h.getAllByCategory(log))
			r.Get("/get-all-by-name", h.getAllByName(log))
			r.Get("/get-all-by-brand", h.getAllByBrand(log))
			r.Put("/update-product", h.updateProduct(log))
			r.Delete("/delete-product", h.deleteProduct(log))
		})

		// ProductImage
		r.Route("/product-image", func(r chi.Router) {
			r.Get("/get-image-by-hash", h.getImageByHash(log))
			r.Get("/get-images-by-product-id", h.getImagesByProductId(log))
			r.Post("/create-product-image", h.createProductImage(log))
			r.Put("/update-product-image", h.updateProductImage(log))
			r.Delete("/delete-product-image-by-name", h.deleteProductImageByName(log))
			r.Delete("/delete-product-image-by-id", h.deleteProductImageById(log))
		})

		r.Route("/cart-content", func(r chi.Router) {
			r.Get("/get-cart-content-by-id", h.getCartContentById(log))
			r.Get("/get-cart-content-by-user-id", h.getCartContentByUserId(log))
			r.Post("/create-cart-content", h.createCartContent(log))
			r.Put("/update-cart-content", h.updateCartContent(log))
			r.Delete("/delete-cart-content", h.deleteCartContent(log))
		})

		//r.Use(middleware.UserIdentity(log, cfg)) // Пример middleware аутентификации
	})

	router.Route("/stream/", func(r chi.Router) {
		r.Route("/chat/", func(r chi.Router) {
			r.Post("/send-message", h.sendMessage(log))
			r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
				websocket.ServeWebSocket(h.hub, w, r, h.services)
			})
		})
	})
}
