package useCases

import (
	"digibank/internal/domain/entities"
	"digibank/internal/interfaceAdapters/dto"
	"digibank/internal/interfaceAdapters/repository"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//go:generate mockgen -source=./transactionUseCase.go -destination=./mocks/transactionUseCase_mock.go
type TransactionUseCase interface {
	CreateTransaction(ctx *gin.Context, transactionContent dto.CreateTransactionInput) (entities.Transaction, error)
}

type transactionUseCase struct {
	repository repository.TransactionRepository
}

func NewTransactionUseCase(repository repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repository}
}

func (a *transactionUseCase) CreateTransaction(ctx *gin.Context, transactionContent dto.CreateTransactionInput) (entities.Transaction, error) {
	transaction := entities.Transaction{
		AccountID:       transactionContent.AccountID,
		OperationTypeID: transactionContent.OperationTypeID,
		Amount:          transactionContent.Amount,
		EventDate:       time.Now(),
	}

	err := a.repository.CreateTransaction(ctx, &transaction)

	if transaction.ID == 0 && err == nil {
		errMsg := fmt.Sprintf("error to create transaction.")
		return transaction, errors.New(errMsg)
	}

	return transaction, err
}
