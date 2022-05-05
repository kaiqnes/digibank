package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckController_HealthCheck(t *testing.T) {
	scenarios := []healthCheckControllerScenario{
		makeScenarioThatGetHealthCheck(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			router := gin.Default()

			testController := NewHealthCheckController(router)
			testController.SetupEndpoints()

			response := httptest.NewRecorder()
			executeRequest(response, scenario.Method, getFullUrl(scenario.Uri, ""), emptyBody, router)

			ctrl.Finish()

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}
