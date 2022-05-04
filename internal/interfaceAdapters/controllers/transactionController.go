package controllers

import (
	"digibank/internal/interfaceAdapters/dto"
	"digibank/internal/interfaceAdapters/presenters"
	"digibank/internal/useCases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transactionsController struct {
	routes    *gin.RouterGroup
	presenter presenters.TransactionPresenter
	useCase   useCases.TransactionUseCase
}

type TransactionsController interface {
	SetupEndpoints()
	transaction(ctx *gin.Context)
}

func NewTransactionsController(routes *gin.RouterGroup, presenter presenters.TransactionPresenter, useCase useCases.TransactionUseCase) TransactionsController {
	return &transactionsController{routes: routes, presenter: presenter, useCase: useCase}
}

func (tc *transactionsController) SetupEndpoints() {
	tc.routes.POST("/transactions", tc.transaction)
}

// transaction 	 godoc
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
func (tc *transactionsController) transaction(ctx *gin.Context) {
	var transactionContent dto.CreateTransactionInput

	if err := ctx.BindJSON(&transactionContent); err != nil {
		tc.presenter.PresentTransactionError(ctx, err, http.StatusBadRequest)
		return
	}

	if _, err := tc.useCase.CreateTransaction(ctx, transactionContent); err != nil {
		// TODO: Create a personal error obj to encapsulate error and a specific error code instead use just error
		tc.presenter.PresentTransactionError(ctx, err, http.StatusInternalServerError)
	} else {
		tc.presenter.PresentTransaction(ctx, http.StatusCreated)
	}
}
