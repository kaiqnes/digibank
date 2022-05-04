package repository

import (
	"digibank/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type accountsRepository struct {
	session *gorm.DB
}

type AccountsRepository interface {
	CreateAccount(ctx *gin.Context, account *entities.Account) error
	GetAccount(ctx *gin.Context, accountID uint) (entities.Account, error)
}

type accountRepository struct {
	session *gorm.DB
}

func NewAccountsRepository(session *gorm.DB) AccountRepository {
	return &accountRepository{session}
}

func (er *accountRepository) CreateAccount(ctx *gin.Context, account *entities.Account) error {
	queryResult := er.session.Create(account)
	return queryResult.Error
}

func (er *accountRepository) GetAccount(ctx *gin.Context, accountID uint) (entities.Account, error) {
	var entityResult entities.Account

	queryResult := er.session.Where("ID", accountID).Find(&entityResult)

	return entityResult, queryResult.Error
}
