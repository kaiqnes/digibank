package presenters

import (
	"github.com/gin-gonic/gin"
)

type transactionPresenter struct {
	errPresenter ErrorPresenter
}

type TransactionPresenter interface {
	PresentTransaction(ctx *gin.Context, statusCode int)
	PresentTransactionError(ctx *gin.Context, err error, statusCode int)
}

type transactionPresenter struct {
	errPresenter ErrorPresenter
}

func NewTransactionPresenter(errPresenter ErrorPresenter) TransactionPresenter {
	return &transactionPresenter{
		errPresenter: errPresenter,
	}
}

func (tp *transactionPresenter) PresentTransaction(ctx *gin.Context, statusCode int) {
	ctx.JSON(statusCode, gin.H{})
}

func (tp *transactionPresenter) PresentTransactionError(ctx *gin.Context, err error, statusCode int) {
	tp.errPresenter.PresentError(ctx, err, statusCode)
}
