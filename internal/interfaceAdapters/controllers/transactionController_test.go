package controllers

import (
	"digibank/internal/interfaceAdapters/presenters"
	mockUseCases "digibank/internal/useCases/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestTransactionController_CreateTransaction(t *testing.T) {
	scenarios := []transactionControllerCreateScenario{
		makeScenarioThatCreateATransaction(),
		makeScenarioThatReceivesABadRequestWhenTryToCreateATransactionSendingANonNumericAccountID(),
		makeScenarioThatReceivesABadRequestWhenTryToCreateATransactionSendingANonNumericOperationTypeID(),
		makeScenarioThatReceivesABadRequestWhenTryToCreateATransactionSendingANonNumericAmount(),
		makeScenarioThatReceivesAnErrorFromUseCaseLayerWhenTryToCreateATransaction(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUseCase := mockUseCases.NewMockTransactionUseCase(ctrl)

			errPresenter := presenters.NewErrorPresenter()
			presenter := presenters.NewTransactionPresenter(errPresenter)
			router := gin.Default()
			group := router.Group(v1)

			testController := NewTransactionController(group, presenter, mockUseCase)
			testController.SetupEndpoints()

			if scenario.ShouldMockUSeCaseCall {
				mockUseCase.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(scenario.MockUseCaseResult, scenario.MockUseCaseError)
			}

			response := httptest.NewRecorder()
			executeRequest(response, scenario.Method, getFullUrl(scenario.Uri, ""), scenario.BodyString, router)

			ctrl.Finish()

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}
