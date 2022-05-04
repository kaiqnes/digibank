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

type transactionUseCase struct {
	repository repository.TransactionsRepository
}

type TransactionUseCase interface {
	CreateTransaction(ctx *gin.Context, transactionContent dto.CreateTransactionInput) (entities.Transaction, error)
}

func NewTransactionUseCase(repository repository.TransactionsRepository) TransactionUseCase {
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
