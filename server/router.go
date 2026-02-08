package server

import (
	"database/sql"
	"kasir-api/internal/handler"
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
	"kasir-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, categoryRepo)
	productHandler := handler.NewProductHandler(productService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	categories := router.Group("/categories")
	{
		categories.GET("/", categoryHandler.GetAll)
		categories.GET("/:id", categoryHandler.GetByID)
		categories.POST("/", categoryHandler.Create)
		categories.PUT("/:id", categoryHandler.Update)
		categories.DELETE("/:id", categoryHandler.Delete)
	}

	products := router.Group("/products")
	{
		products.GET("/", productHandler.GetAll)
		products.GET("/:id", productHandler.GetByID)
		products.POST("/", productHandler.Create)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)
	}

	api := router.Group("/api")
	{
		api.POST("/checkout", transactionHandler.Checkout)
	}

	router.GET("/health", func(c *gin.Context) {
		resp := model.Response{
			Message: "OK",
		}
		c.JSON(http.StatusOK, resp)
	})

	return router
}
