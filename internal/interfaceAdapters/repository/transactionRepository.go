package repository

import (
	"digibank/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//go:generate mockgen -source=./transactionRepository.go -destination=./mocks/transactionRepository_mock.go
type TransactionRepository interface {
	CreateTransaction(ctx *gin.Context, transaction *entities.Transaction) error
}

type transactionRepository struct {
	session *gorm.DB
}

func NewTransactionRepository(session *gorm.DB) TransactionRepository {
	return &transactionRepository{session}
}

func (er *transactionRepository) CreateTransaction(ctx *gin.Context, transaction *entities.Transaction) error {
	queryResult := er.session.Create(transaction)
	return queryResult.Error
}
