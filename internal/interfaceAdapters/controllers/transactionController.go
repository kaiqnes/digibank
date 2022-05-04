package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type transactionsController struct {
	routes *gin.RouterGroup
}

type TransactionsController interface {
	SetupEndpoints()
	transaction(ctx *gin.Context)
}

func NewTransactionsController(routes *gin.RouterGroup, db *gorm.DB) TransactionsController {
	return &transactionsController{routes: routes}
}

func (h *transactionsController) SetupEndpoints() {
	h.routes.POST("/transactions", h.transaction)
}

// transaction 	 godoc
// @Summary      This endpoint receives a transaction
// @Description  This endpoint receives a transaction
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /transactions [post]
func (h *transactionsController) transaction(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}
