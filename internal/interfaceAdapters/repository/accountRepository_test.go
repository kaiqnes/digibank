package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func TestAccountRepository_CreateAccount(t *testing.T) {
	scenarios := []accountRepositoryScenario{
		makeScenarioThatCreatesAnAccount(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			session := beforeEach()
			repository := NewAccountsRepository(session)

			if len(scenario.MockRows) > 0 {
				for _, row := range scenario.MockRows {
					session.Create(&row)
				}
			}

			testContext, _ := gin.CreateTestContext(httptest.NewRecorder())
			err := repository.CreateAccount(testContext, &scenario.Data)

			assert.Equal(t, err, scenario.ExpectError)
			assert.Equal(t, scenario.Data.ID, scenario.ExpectID)
		})
	}

	removeTestDB()
}

func TestAccountRepository_GetAccount(t *testing.T) {
	scenarios := []accountRepositoryScenario{
		makeScenarioThatRetrievesAnAccount(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			session := beforeEach()
			repository := NewAccountsRepository(session)

			if len(scenario.MockRows) > 0 {
				for _, row := range scenario.MockRows {
					session.Create(&row)
				}
			}

			testContext, _ := gin.CreateTestContext(httptest.NewRecorder())
			account, err := repository.GetAccount(testContext, scenario.Data.ID)

			assert.Equal(t, account.DocumentNumber, scenario.ExpectDocumentNumber)
			assert.Equal(t, account.ID, scenario.ExpectID)
			assert.Equal(t, err, scenario.ExpectError)
		})
	}

	removeTestDB()
}
