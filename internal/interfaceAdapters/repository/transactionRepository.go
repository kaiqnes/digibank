package repository

import (
	"digibank/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type transactionsRepository struct {
	session *gorm.DB
}

type TransactionsRepository interface {
	CreateTransaction(ctx *gin.Context, transaction *entities.Transaction) error
}

func NewTransactionsRepository(session *gorm.DB) TransactionsRepository {
	return &transactionsRepository{session}
}

func (er *transactionsRepository) CreateTransaction(ctx *gin.Context, transaction *entities.Transaction) error {
	queryResult := er.session.Create(transaction)
	return queryResult.Error
}
