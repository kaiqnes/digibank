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

type accountController struct {
	routes    *gin.RouterGroup
	presenter presenters.AccountPresenter
	useCase   useCases.AccountUseCase
}

type AccountController interface {
	SetupEndpoints()
	createAccount(ctx *gin.Context)
	getAccount(ctx *gin.Context)
}

func NewAccountController(routes *gin.RouterGroup, presenter presenters.AccountPresenter, useCase useCases.AccountUseCase) AccountController {
	return &accountController{routes: routes, presenter: presenter, useCase: useCase}
}

func (a *accountController) SetupEndpoints() {
	a.routes.POST("/accounts", a.createAccount)
	a.routes.GET("/accounts/:accountID", a.getAccount)
}

// createAccount 	 godoc
// @Summary      This endpoint creates a new account
// @Description  This endpoint creates a new account and returns the "document_number" (string) representing the accountID
// @Accept       json
// @Produce      json
// @Param request_body body string true "Document number to be inserted into account to be created" SchemaExample({\r\n  "document_number": "12345678900"\r\n})
// @Success      200 {object} dto.AccountPresenterResponse
// @Failure      400 {object} dto.ErrorOutputDto
// @Failure      500 {object} dto.ErrorOutputDto
// @Router       /api/v1/accounts [post]
func (a *accountController) createAccount(ctx *gin.Context) {
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

// getAccount 	 godoc
// @Summary      This endpoint retrieves a specific account
// @Description  This endpoint receives an accountID and returns the respective account details
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param 		 accountID path int true "Account ID"
// @Success      200 {object} dto.AccountPresenterResponse
// @Failure      400 {object} dto.ErrorOutputDto
// @Failure      500 {object} dto.ErrorOutputDto
// @Router       /api/v1/accounts/{accountID} [get]
func (a *accountController) getAccount(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	uAccountID, err := validateAccountID(accountID)
	if err != nil {
		a.presenter.PresentAccountError(ctx, err, http.StatusBadRequest)
		return
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
