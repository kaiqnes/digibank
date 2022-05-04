package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type accountsController struct {
	routes *gin.RouterGroup
	db     *gorm.DB
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
	ctx.JSON(http.StatusOK, gin.H{
		"mock_document_number": "12345678900",
	})
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
	ctx.JSON(http.StatusOK, gin.H{
		"mock_account_id":      "1",
		"mock_document_number": "12345678900",
	})
}

type AccountsController interface {
	SetupEndpoints()
	createAccount(ctx *gin.Context)
	getAccount(ctx *gin.Context)
}

func NewAccountsController(routes *gin.RouterGroup, db *gorm.DB) AccountsController {
	return &accountsController{routes: routes, db: db}
}
