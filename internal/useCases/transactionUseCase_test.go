package useCases

import (
	mockRepository "digibank/internal/interfaceAdapters/repository/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestTransactionUseCase_CreateTransaction(t *testing.T) {
	scenarios := []transactionUseCaseCreateScenario{
		makeScenarioThatCreateATransaction(),
		makeScenarioThatReceivesAnErrorWhenCreateATransaction(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := mockRepository.NewMockTransactionRepository(ctrl)
			testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

			mockRepo.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(scenario.MockRepoResult, scenario.MockRepoError)

			useCase := NewTransactionUseCase(mockRepo)

			result, err := useCase.CreateTransaction(testContext, scenario.Data)

			assert.Equal(t, err, scenario.ExpectError)
			assert.Equal(t, result.OperationTypeID, scenario.ExpectResult.OperationTypeID)
			assert.Equal(t, result.AccountID, scenario.ExpectResult.AccountID)
			assert.Equal(t, result.Amount, scenario.ExpectResult.Amount)
		})
	}
}
