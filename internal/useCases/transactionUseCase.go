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
	"time"
)

//go:generate mockgen -source=./transactionUseCase.go -destination=./mocks/transactionUseCase_mock.go
type TransactionUseCase interface {
	CreateTransaction(ctx *gin.Context, transactionContent dto.CreateTransactionInput) (entities.Transaction, errorx.Errorx)
}

type transactionUseCase struct {
	repository repository.TransactionRepository
}

func NewTransactionUseCase(repository repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repository}
}

func (a *transactionUseCase) CreateTransaction(ctx *gin.Context, transactionContent dto.CreateTransactionInput) (entities.Transaction, errorx.Errorx) {
	var errx errorx.Errorx

	transaction := entities.Transaction{
		AccountID:       transactionContent.AccountID,
		OperationTypeID: transactionContent.OperationTypeID,
		Amount:          transactionContent.Amount,
		EventDate:       time.Now(),
	}

	createdTransaction, err := a.repository.CreateTransaction(ctx, transaction)

	if createdTransaction.ID == 0 && err == nil {
		errMsg := fmt.Sprintf("error to create transaction.")
		return createdTransaction, errorx.NewErrorx(http.StatusInternalServerError, errors.New(errMsg))
	}

	if err != nil {
		errx = errorx.NewErrorx(http.StatusInternalServerError, err)
	}

	return createdTransaction, errx
}
