package controllers

import (
	"digibank/internal/interfaceAdapters/dto"
	"digibank/internal/interfaceAdapters/presenters"
	"digibank/internal/useCases"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type accountsController struct {
	routes    *gin.RouterGroup
	presenter presenters.AccountPresenter
	useCase   useCases.AccountsUseCase
}

type AccountsController interface {
	SetupEndpoints()
	createAccount(ctx *gin.Context)
	getAccount(ctx *gin.Context)
}

func NewAccountsController(routes *gin.RouterGroup, presenter presenters.AccountPresenter, useCase useCases.AccountsUseCase) AccountsController {
	return &accountsController{routes: routes, presenter: presenter, useCase: useCase}
}

func (a *accountsController) SetupEndpoints() {
	a.routes.POST("/accounts", a.createAccount)
	a.routes.GET("/accounts/:accountID", a.getAccount)
}

// createAccount 	 godoc
// @Summary      This endpoint creates a new account
// @Description  This endpoint creates a new account and returns the "document_number" (string) representing the accountID
// @Tags         CreateAccount
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /accounts [post]
func (a *accountsController) createAccount(ctx *gin.Context) {
	var accountContent dto.CreateAccountInput

	if err := ctx.BindJSON(&accountContent); err != nil {
		a.presenter.PresentAccountError(ctx, err, http.StatusBadRequest)
		return
	}

	if createdAccount, err := a.useCase.CreateAccount(ctx, accountContent); err != nil {
		// TODO: Create a personal error obj to encapsulate error and a specific error code instead use just error
		a.presenter.PresentAccountError(ctx, err, http.StatusInternalServerError)
	} else {
		a.presenter.PresentAccount(ctx, createdAccount, http.StatusCreated)
	}
}

// createAccount 	 godoc
// @Summary      This endpoint retrieves a specific account
// @Description  This endpoint receives an accountID and returns the respective account details
// @Tags         GetAccount
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /accounts/:accountID [get]
func (a *accountsController) getAccount(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, gin.H{
	//	"mock_account_id":      "1",
	//	"mock_document_number": "12345678900",
	//})
	accountID := ctx.Param("accountID")
	uAccountID, err := validateAccountID(accountID)
	if err != nil {
		a.presenter.PresentAccountError(ctx, err, http.StatusBadRequest)
	}

	if account, err := a.useCase.GetAccount(ctx, uAccountID); err != nil {
		a.presenter.PresentAccountError(ctx, err, http.StatusInternalServerError)
	} else {
		a.presenter.PresentAccount(ctx, account, http.StatusOK)
	}
}

func validateAccountID(id string) (uint, error) {
	if id == "" {
		errMsg := fmt.Sprintf("invalid accountID received: '%s'. accountID must be not null, numeric and bigger than zero.", id)
		return 0, errors.New(errMsg)
	}

	uID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		errMsg := fmt.Sprintf("invalid accountID received: '%s'. accountID must be not null, numeric and bigger than zero.", id)
		return 0, errors.New(errMsg)
	}

	if uID < 1 {
		errMsg := fmt.Sprintf("invalid accountID received: '%s'. accountID must be not null, numeric and bigger than zero.", id)
		return 0, errors.New(errMsg)
	}

	return uint(uID), nil
}
