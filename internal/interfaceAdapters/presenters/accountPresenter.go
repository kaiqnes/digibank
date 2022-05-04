package presenters

import (
	"digibank/internal/domain/entities"
	"digibank/internal/interfaceAdapters/dto"
	"github.com/gin-gonic/gin"
)

type accountPresenter struct{}

type AccountPresenter interface {
	PresentAccount(ctx *gin.Context, account entities.Account, statusCode int)
	PresentAccountError(ctx *gin.Context, err error, statusCode int)
}

func NewAccountPresenter() AccountPresenter {
	return &accountPresenter{}
}

func (pbl *accountPresenter) PresentAccount(ctx *gin.Context, account entities.Account, statusCode int) {
	ctx.JSON(statusCode, pbl.generateAccountResponse(account))
}

func (pbl *accountPresenter) PresentAccountError(ctx *gin.Context, err error, statusCode int) {
	errResponse := dto.ErrorOutputDto{Message: err.Error()}
	ctx.JSON(statusCode, errResponse)
}

func (pbl *accountPresenter) generateAccountResponse(account entities.Account) dto.AccountPresenterResponse {
	var response dto.AccountPresenterResponse

	response.AccountID = account.ID
	response.DocumentNumber = account.DocumentNumber

	return response
}
