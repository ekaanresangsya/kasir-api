package handler

import (
	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "success",
		Data:    categories,
	})
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid id",
		})
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "success",
		Data:    category,
	})
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category *model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid request body",
		})
		return
	}

	category, err := h.service.Create(category)
	if err != nil {
		log.Printf("error creating category, got %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "success",
		Data:    category,
	})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid id",
		})
		return
	}

	var category *model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid request body",
		})
		return
	}

	category.ID = id
	err = h.service.Update(id, category)
	if err != nil {
		log.Printf("error updating category, got %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "success",
		Data:    category,
	})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "invalid id",
		})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "success",
	})
}
