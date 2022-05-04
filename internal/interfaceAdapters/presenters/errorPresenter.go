package presenters

import (
	"digibank/internal/interfaceAdapters/dto"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./errorPresenter.go -destination=./mocks/errorPresenter_mock.go
type ErrorPresenter interface {
	PresentError(ctx *gin.Context, err error, statusCode int)
}

type errorPresenter struct{}

func NewErrorPresenter() ErrorPresenter {
	return &errorPresenter{}
}

func (pbl *errorPresenter) PresentError(ctx *gin.Context, err error, statusCode int) {
	errResponse := dto.ErrorOutputDto{Message: err.Error()}
	ctx.JSON(statusCode, errResponse)
}
