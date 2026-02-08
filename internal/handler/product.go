package handler

import (
	"fmt"
	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetAll(c *gin.Context) {

	req := model.GetProductReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}

	fmt.Println(req)

	products, err := h.productService.GetAll(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	productResponses := make([]*model.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = product.ToResponse()
	}

	resp := model.Response{
		Message: "success",
		Data:    productResponses,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: "Invalid product ID"})
		return
	}

	product, err := h.productService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	resp := model.Response{
		Message: "success",
		Data:    product.ToResponse(),
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req model.CreateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}

	product := &model.Product{
		Name:       req.Name,
		Stock:      req.Stock,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	}

	err := h.productService.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "success",
		Data:    product.ToResponse(),
	})
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: "Invalid product ID"})
		return
	}

	var req model.UpdateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}

	product := &model.Product{
		Name:       req.Name,
		Stock:      req.Stock,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	}

	product.ID = int64(id)
	if err := h.productService.Update(product); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "Product updated successfully",
		Data:    product.ToResponse(),
	})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: "Invalid product ID"})
		return
	}

	if err := h.productService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{Message: "Product deleted successfully"})
}
