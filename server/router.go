package server

import (
	"crud-categories/internal/handler"
	"crud-categories/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/categories", handler.GetAllCategories)
	router.GET("/categories/:id", handler.GetCategoryByID)
	router.POST("/categories", handler.CreateCategory)
	router.PUT("/categories/:id", handler.UpdateCategory)
	router.DELETE("/categories/:id", handler.DeleteCategory)

	router.GET("/health", func(c *gin.Context) {
		resp := model.Response{
			Message: "OK",
		}
		c.JSON(http.StatusOK, resp)
	})

	return router
}
