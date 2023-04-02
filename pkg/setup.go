package pkg

import (
	"fmt"

	"github.com/APoniatowski/listingd/config"
	"github.com/APoniatowski/listingd/pkg/handlers"
	"github.com/APoniatowski/listingd/pkg/middlewares"
	"github.com/APoniatowski/listingd/repositories"
	"github.com/APoniatowski/listingd/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(cfg *config.Config) *gorm.DB {
	// Setting up Data Source Name
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
		cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode,
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize handlers
	companyRepo := repositories.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo)
	companyHandler := handlers.NewCompanyHandler(companyService)

	// Initialize Gin router
	router := gin.Default()

	// Initialize the JWT middleware
	authMiddleware, err := middlewares.JWTAuthMiddleware()
	if err != nil {
		panic("Error initializing JWT middleware")
	}

	// Public routes
	router.GET("/companies/:id", companyHandler.GetCompany)
	router.POST("/login", authMiddleware.LoginHandler)

	// Use the JWT middleware for protected routes
	authorized := router.Group("/")
	authorized.Use(authMiddleware.MiddlewareFunc())
	{
		authorized.POST("/companies", companyHandler.CreateCompany)
		authorized.PATCH("/companies/:id", companyHandler.PatchCompany)
		authorized.DELETE("/companies/:id", companyHandler.DeleteCompany)
	}

	return router
}
