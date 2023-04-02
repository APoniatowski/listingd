package tests

import (
	"log"

	"github.com/APoniatowski/listingd/config"
	"github.com/APoniatowski/listingd/pkg"
	"github.com/APoniatowski/listingd/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	router *gin.Engine
	db     *gorm.DB
)

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db = pkg.SetupDatabase(cfg)
	router = pkg.SetupRouter(db)

	err = db.AutoMigrate(&models.Company{})
	if err != nil {
		panic("Failed to migrate the test database")
	}
}
