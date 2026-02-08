package handler

import (
	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) Checkout(c *gin.Context) {
	var req model.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
		return
	}

	transaction, err := h.transactionService.Checkout(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	resData := model.Response{
		Message: "Checkout successful",
		Data:    transaction,
	}

	c.JSON(http.StatusOK, resData)
}
