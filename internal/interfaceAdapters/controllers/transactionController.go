package controllers

import (
	"digibank/internal/interfaceAdapters/dto"
	"digibank/internal/interfaceAdapters/presenters"
	"digibank/internal/useCases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transactionController struct {
	routes    *gin.RouterGroup
	presenter presenters.TransactionPresenter
	useCase   useCases.TransactionUseCase
}

type TransactionController interface {
	SetupEndpoints()
	createTransaction(ctx *gin.Context)
}

func NewTransactionController(routes *gin.RouterGroup, presenter presenters.TransactionPresenter, useCase useCases.TransactionUseCase) TransactionController {
	return &transactionController{routes: routes, presenter: presenter, useCase: useCase}
}

func (tc *transactionController) SetupEndpoints() {
	tc.routes.POST("/transactions", tc.createTransaction)
}

// createTransaction 	 godoc
// @Summary      This endpoint receives a transaction
// @Description  This endpoint receives a transaction
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param request_body body string true "Body of a transaction" SchemaExample({\r\n"account_id": 1,\r\n"operation_type_id": 4,\r\n"amount": 123.45\r\n})
// @Success      200 {object} interface{}
// @Failure      400 {object} dto.ErrorOutputDto
// @Failure      500 {object} dto.ErrorOutputDto
// @Router       /api/v1/transactions [post]
func (tc *transactionController) createTransaction(ctx *gin.Context) {
	var transactionContent dto.CreateTransactionInput

	if err := ctx.BindJSON(&transactionContent); err != nil {
		tc.presenter.PresentTransactionError(ctx, err, http.StatusBadRequest)
		return
	}

	if _, errx := tc.useCase.CreateTransaction(ctx, transactionContent); errx != nil {
		tc.presenter.PresentTransactionError(ctx, errx.GetError(), errx.GetStatusCode())
	} else {
		tc.presenter.PresentTransaction(ctx, http.StatusCreated)
	}
}
