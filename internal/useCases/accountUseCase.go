package useCases

import (
	"digibank/internal/domain/entities"
	"digibank/internal/interfaceAdapters/dto"
	"digibank/internal/interfaceAdapters/repository"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type accountsUseCase struct {
	repository repository.AccountsRepository
}

type AccountsUseCase interface {
	CreateAccount(ctx *gin.Context, accountContent dto.CreateAccountInput) (entities.Account, error)
	GetAccount(ctx *gin.Context, accountID uint) (entities.Account, error)
}

type accountUseCase struct {
	repository repository.AccountRepository
}

func NewAccountUseCase(repository repository.AccountRepository) AccountUseCase {
	return &accountUseCase{repository}
}

func (a *accountUseCase) CreateAccount(ctx *gin.Context, accountContent dto.CreateAccountInput) (entities.Account, error) {
	account := entities.Account{DocumentNumber: accountContent.DocumentNumber}
	err := a.repository.CreateAccount(ctx, &account)
	return account, err
}

func (a *accountUseCase) GetAccount(ctx *gin.Context, accountID uint) (entities.Account, error) {
	account, err := a.repository.GetAccount(ctx, accountID)

	if account.ID == 0 && err == nil {
		errMsg := fmt.Sprintf("account with id %d wasn't found in db.", accountID)
		return account, errors.New(errMsg)
	}

	return account, err
}
