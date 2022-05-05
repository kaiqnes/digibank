package useCases

import (
	mockRepository "digibank/internal/interfaceAdapters/repository/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestAccountUseCase_CreateAccount(t *testing.T) {
	scenarios := []accountUseCaseCreateScenario{
		makeScenarioThatCreateAnAccount(),
		makeScenarioThatReceivesAnErrorWhenCreateAnAccount(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := mockRepository.NewMockAccountRepository(ctrl)
			testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

			mockRepo.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(scenario.MockRepoResult, scenario.MockRepoError)

			useCase := NewAccountUseCase(mockRepo)

			result, err := useCase.CreateAccount(testContext, scenario.Data)

			assert.Equal(t, err, scenario.ExpectError)
			assert.Equal(t, result.DocumentNumber, scenario.ExpectDocumentNumber)
		})
	}
}

func TestAccountUseCase_GetAccount(t *testing.T) {
	scenarios := []accountUseCaseGetScenario{
		makeScenarioThatRetrievesAnAccount(),
		makeScenarioThatRetrievesNoneAccount(),
		makeScenarioThatReceivesAnErrorWhenGetAnAccount(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := mockRepository.NewMockAccountRepository(ctrl)
			testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

			mockRepo.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Return(scenario.MockRepoResult, scenario.MockRepoError)

			useCase := NewAccountUseCase(mockRepo)

			result, err := useCase.GetAccount(testContext, scenario.AccountID)

			assert.Equal(t, err, scenario.ExpectError)
			assert.Equal(t, result.DocumentNumber, scenario.ExpectDocumentNumber)
		})
	}
}
