package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	scenarios := []transactionRepositoryScenario{
		makeScenarioThatCreatesATransaction(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			session := beforeEach()
			repository := NewTransactionRepository(session)

			if len(scenario.MockRows) > 0 {
				for _, row := range scenario.MockRows {
					session.Create(&row)
				}
			}

			testContext, _ := gin.CreateTestContext(httptest.NewRecorder())
			err := repository.CreateTransaction(testContext, &scenario.Data)

			assert.Equal(t, err, scenario.ExpectError)
			assert.Equal(t, scenario.Data.ID, scenario.ExpectID)
		})
	}

	removeTestDB()
}
