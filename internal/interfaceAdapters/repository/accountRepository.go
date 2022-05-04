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

func NewAccountsRepository(session *gorm.DB) AccountsRepository {
	return &accountsRepository{session}
}

func (er *accountsRepository) CreateAccount(ctx *gin.Context, account *entities.Account) error {
	queryResult := er.session.Create(account)
	return queryResult.Error
}

func (er *accountsRepository) GetAccount(ctx *gin.Context, accountID uint) (entities.Account, error) {
	var entityResult entities.Account

	queryResult := er.session.Where("ID", accountID).Find(&entityResult)

	return entityResult, queryResult.Error
}
