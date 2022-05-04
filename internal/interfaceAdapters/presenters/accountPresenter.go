package presenters

import (
	"digibank/internal/domain/entities"
	"digibank/internal/interfaceAdapters/dto"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./accountPresenter.go -destination=./mocks/accountPresenter_mock.go
type AccountPresenter interface {
	PresentAccount(ctx *gin.Context, account entities.Account, statusCode int)
	PresentAccountError(ctx *gin.Context, err error, statusCode int)
}

type accountPresenter struct {
	errPresenter ErrorPresenter
}

func NewAccountPresenter(errPresenter ErrorPresenter) AccountPresenter {
	return &accountPresenter{
		errPresenter: errPresenter,
	}
}

func (ap *accountPresenter) PresentAccount(ctx *gin.Context, account entities.Account, statusCode int) {
	ctx.JSON(statusCode, ap.generateAccountResponse(account))
}

func (ap *accountPresenter) PresentAccountError(ctx *gin.Context, err error, statusCode int) {
	ap.errPresenter.PresentError(ctx, err, statusCode)
}

func (ap *accountPresenter) generateAccountResponse(account entities.Account) dto.AccountPresenterResponse {
	var response dto.AccountPresenterResponse

	response.AccountID = account.ID
	response.DocumentNumber = account.DocumentNumber

	return response
}
