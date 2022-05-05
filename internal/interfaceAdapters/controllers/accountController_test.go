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

func TestAccountController_CreateAccount(t *testing.T) {
	scenarios := []accountControllerCreateScenario{
		makeScenarioThatCreateAnAccount(),
		makeScenarioThatReceivesABadRequestWhenTryToCreateAnAccount(),
		makeScenarioThatReceivesAnErrorFromUseCaseLayerWhenTryToCreateAnAccount(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUseCase := mockUseCases.NewMockAccountUseCase(ctrl)

			errPresenter := presenters.NewErrorPresenter()
			presenter := presenters.NewAccountPresenter(errPresenter)
			router := gin.Default()
			group := router.Group(v1)

			testController := NewAccountController(group, presenter, mockUseCase)
			testController.SetupEndpoints()

			if scenario.ShouldMockUSeCaseCall {
				mockUseCase.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(scenario.MockUseCaseResult, scenario.MockUseCaseError)
			}

			response := httptest.NewRecorder()
			executeRequest(response, scenario.Method, getFullUrl(scenario.Uri, ""), scenario.BodyString, router)

			ctrl.Finish()

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}

func TestAccountController_GetAccount(t *testing.T) {
	scenarios := []accountControllerGetScenario{
		makeScenarioThatGetAnAccount(),
		makeScenarioThatReceivesABadRequestWhenTryToGetAnAccountSendingANegativeAccountID(),
		makeScenarioThatReceivesABadRequestWhenTryToGetAnAccountSendingANonNumericAccountID(),
		makeScenarioThatReceivesAnErrorFromUseCaseLayerWhenTryToGetAnAccount(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUseCase := mockUseCases.NewMockAccountUseCase(ctrl)

			errPresenter := presenters.NewErrorPresenter()
			presenter := presenters.NewAccountPresenter(errPresenter)
			router := gin.Default()
			group := router.Group(v1)

			testController := NewAccountController(group, presenter, mockUseCase)
			testController.SetupEndpoints()

			if scenario.ShouldMockUSeCaseCall {
				mockUseCase.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Return(scenario.MockUseCaseResult, scenario.MockUseCaseError)
			}

			response := httptest.NewRecorder()
			executeRequest(response, scenario.Method, getFullUrl(scenario.Uri, scenario.PathParam), scenario.BodyString, router)

			ctrl.Finish()

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}
