package presenters

import (
	"digibank/internal/interfaceAdapters/dto"
	"github.com/gin-gonic/gin"
)

type errorPresenter struct{}

type ErrorPresenter interface {
	PresentError(ctx *gin.Context, err error, statusCode int)
}

func NewErrorPresenter() ErrorPresenter {
	return &errorPresenter{}
}

func (pbl *errorPresenter) PresentError(ctx *gin.Context, err error, statusCode int) {
	errResponse := dto.ErrorOutputDto{Message: err.Error()}
	ctx.JSON(statusCode, errResponse)
}
