package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/config"
	"github.com/kunalavghade/Go-basics/go-api/internal/interfaces"
	"github.com/kunalavghade/Go-basics/go-api/internal/providers"
	"github.com/kunalavghade/Go-basics/go-api/internal/services"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	DB             *gorm.DB
	log            *zerolog.Logger
	config         *config.Config
	authService    *services.AuthService
	userService    *services.UserService
	productService *services.ProductService
	uploadService  *services.UploadService
}

func NewServer(cfg *config.Config, log *zerolog.Logger, db *gorm.DB) *Server {
	var uploadProvider interfaces.UploadProvider
	if cfg.Upload.Provider == "aws" {
		uploadProvider = providers.NewS3Provider(cfg)
	} else {
		uploadProvider = providers.NewLocalUploadProvider(cfg.Upload.Path)
	}
	return &Server{
		DB:             db,
		log:            log,
		config:         cfg,
		authService:    services.NewAuthService(db, cfg),
		userService:    services.NewUserService(db),
		productService: services.NewProductService(db),
		uploadService:  services.NewUploadService(uploadProvider),
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	// Add Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	// Add Routes
	router.GET("/health", s.HealthCheck)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{ // nolint:gocritic // Manage readbility
			auth.POST("/register", s.Register)
			auth.POST("/login", s.Login)
			auth.POST("/refresh", s.RefreshToken)
			auth.POST("/logout", s.Logout)
		}

		protected := api.Group("/")
		protected.Use(s.authMiddleware())
		{ // nolint:gocritic // Manage readbility
			// Users
			users := protected.Group("/users")
			users.GET("/profile", s.GetProfile)
			users.PUT("/profile", s.UpdateProfile)

			// Products
			products := protected.Group("/products")
			products.POST("/", s.CreateProductHandler)
			products.GET("/", s.GetProductsHandler)
			products.GET("/:id", s.GetProductByIDHandler)
			products.PUT("/:id", s.UpdateProductHandler)
			products.DELETE("/:id", s.DeleteProductHandler)
			products.POST("/:id/image", s.UploadProductImageHandler)

			// Categories
			categories := protected.Group("/categories")
			categories.POST("/", s.CreateCategoryHandler)
			categories.GET("/", s.GetCategoriesHandler)
			categories.PUT("/:id", s.UpdateCategoryHandler)
			categories.DELETE("/:id", s.DeleteCategoryHandler)
		}

	}

	return router
}

func (s *Server) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
