package main

import (
	_ "backend/docs"
	"backend/internal/config"
	"backend/internal/lib/logger/handlers/slogpretty"
	"backend/internal/middleware/logger"
	"backend/internal/repository"
	"backend/internal/routes"
	"backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	_ "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	//envDev   = "dev"
	envProd = "prod"
)

// @title           Stream Service Prototype API
// @version         1.0
// @description     API server for stream service application

// @host      localhost:8082
// @BasePath  /

// @securityDefinitions.apiKey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password, // TODO: убрать в env
	})
	if err != nil {
		log.Error("cannot connect to database:", slog.String("error", err.Error()))
		os.Exit(1)
	} else {
		log.Info("connected to database")
	}

	repos := repository.NewRepositories(db)
	services := service.NewService(repos)
	handlers := routes.NewHandler(services)

	router := chi.NewRouter()

	// Настройка CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   cfg.CorsOrigin, // Разрешаем фронтенд на порту 3000
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	// Применяем CORS middleware
	router.Use(corsMiddleware.Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	handlers.RegisterRoutes(router, log, &cfg)

	log.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if cfg.Env == envProd {
		if err := srv.ListenAndServeTLS(cfg.HTTPServer.CertFile, cfg.HTTPServer.KeyFile); err != nil {
			log.Error("cannot start server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	} else {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("cannot start server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}

	log.Info("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
