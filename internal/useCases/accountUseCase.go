package useCases

import (
	"digibank/internal/domain/entities"
	"digibank/internal/frameworks/errorx"
	"digibank/internal/interfaceAdapters/dto"
	"digibank/internal/interfaceAdapters/repository"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate mockgen -source=./accountUseCase.go -destination=./mocks/accountUseCase_mock.go
type AccountUseCase interface {
	CreateAccount(ctx *gin.Context, accountContent dto.CreateAccountInput) (entities.Account, errorx.Errorx)
	GetAccount(ctx *gin.Context, accountID uint) (entities.Account, errorx.Errorx)
}

type accountUseCase struct {
	repository repository.AccountRepository
}

func NewAccountUseCase(repository repository.AccountRepository) AccountUseCase {
	return &accountUseCase{repository}
}

func (a *accountUseCase) CreateAccount(ctx *gin.Context, accountContent dto.CreateAccountInput) (entities.Account, errorx.Errorx) {
	var errx errorx.Errorx

	createdAccount, err := a.repository.CreateAccount(ctx, entities.Account{DocumentNumber: accountContent.DocumentNumber})

	if err != nil {
		errx = errorx.NewErrorx(http.StatusInternalServerError, err)
	}

	return createdAccount, errx
}

func (a *accountUseCase) GetAccount(ctx *gin.Context, accountID uint) (entities.Account, errorx.Errorx) {
	var errx errorx.Errorx

	account, err := a.repository.GetAccount(ctx, accountID)

	if account.ID == 0 && err == nil {
		errMsg := fmt.Sprintf("account with id %d wasn't found in db.", accountID)
		return account, errorx.NewErrorx(http.StatusInternalServerError, errors.New(errMsg))
	}

	if err != nil {
		errx = errorx.NewErrorx(http.StatusInternalServerError, err)
	}

	return account, errx
}
