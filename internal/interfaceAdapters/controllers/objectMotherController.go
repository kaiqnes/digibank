package controllers

import (
	"bytes"
	"digibank/internal/domain/entities"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
)

const (
	baseUrl        = "http://localhost:8080"
	v1             = "/api/v1"
	accountUri     = "/api/v1/accounts"
	transactionUri = "/api/v1/transactions"
	emptyBody      = ""
)

type accountControllerCreateScenario struct {
	TestName              string
	Method                string
	Uri                   string
	BodyString            string
	ShouldMockUSeCaseCall bool
	MockUseCaseResult     entities.Account
	MockUseCaseError      error
	ExpectStatus          int
	ExpectResponse        string
}

type accountControllerGetScenario struct {
	TestName              string
	Method                string
	Uri                   string
	PathParam             string
	BodyString            string
	ShouldMockUSeCaseCall bool
	MockUseCaseResult     entities.Account
	MockUseCaseError      error
	ExpectStatus          int
	ExpectResponse        string
}

type transactionControllerCreateScenario struct {
	TestName              string
	Method                string
	Uri                   string
	BodyString            string
	ShouldMockUSeCaseCall bool
	MockUseCaseResult     entities.Transaction
	MockUseCaseError      error
	ExpectStatus          int
	ExpectResponse        string
}

func makeScenarioThatCreateAnAccount() accountControllerCreateScenario {
	return accountControllerCreateScenario{
		TestName:              "Create an account",
		Method:                http.MethodPost,
		Uri:                   accountUri,
		BodyString:            `{"document_number": "12345678900"}`,
		ShouldMockUSeCaseCall: true,
		MockUseCaseResult:     entities.Account{Model: gorm.Model{ID: 1}, DocumentNumber: "12345678900"},
		MockUseCaseError:      nil,
		ExpectStatus:          http.StatusCreated,
		ExpectResponse:        `{"account_id":1,"document_number":"12345678900"}`,
	}
}

func makeScenarioThatReceivesABadRequestWhenTryToCreateAnAccount() accountControllerCreateScenario {
	return accountControllerCreateScenario{
		TestName:              "Receives a bad request when send in 'document_number' a non-string value",
		Method:                http.MethodPost,
		Uri:                   accountUri,
		BodyString:            `{"document_number": 1234}`,
		ShouldMockUSeCaseCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"message":"json: cannot unmarshal number into Go struct field CreateAccountInput.document_number of type string"}`,
	}
}

func makeScenarioThatReceivesAnErrorFromUseCaseLayerWhenTryToCreateAnAccount() accountControllerCreateScenario {
	return accountControllerCreateScenario{
		TestName:              "Receives an error from use case layer",
		Method:                http.MethodPost,
		Uri:                   accountUri,
		BodyString:            `{"document_number": "12345678900"}`,
		ShouldMockUSeCaseCall: true,
		MockUseCaseResult:     entities.Account{},
		MockUseCaseError:      errors.New("mock-error"),
		ExpectStatus:          http.StatusInternalServerError,
		ExpectResponse:        `{"message":"mock-error"}`,
	}
}

func makeScenarioThatGetAnAccount() accountControllerGetScenario {
	return accountControllerGetScenario{
		TestName:              "Get an account",
		Method:                http.MethodGet,
		Uri:                   accountUri,
		PathParam:             "1",
		BodyString:            emptyBody,
		ShouldMockUSeCaseCall: true,
		MockUseCaseResult:     entities.Account{Model: gorm.Model{ID: 1}, DocumentNumber: "12345678900"},
		ExpectStatus:          http.StatusOK,
		ExpectResponse:        `{"account_id":1,"document_number":"12345678900"}`,
	}
}

func makeScenarioThatReceivesABadRequestWhenTryToGetAnAccountSendingANegativeAccountID() accountControllerGetScenario {
	return accountControllerGetScenario{
		TestName:              "Receives a bad request when send in 'accountID path param' a negative numeric value",
		Method:                http.MethodGet,
		Uri:                   accountUri,
		PathParam:             "-1",
		BodyString:            emptyBody,
		ShouldMockUSeCaseCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"message":"invalid accountID received: '-1'. accountID must be not null, numeric and bigger than zero."}`,
	}
}

func makeScenarioThatReceivesABadRequestWhenTryToGetAnAccountSendingANonNumericAccountID() accountControllerGetScenario {
	return accountControllerGetScenario{
		TestName:              "Receives a not_found when send in 'accountID path param' a non-numeric value",
		Method:                http.MethodGet,
		Uri:                   accountUri,
		PathParam:             "invalid_account_id",
		BodyString:            emptyBody,
		ShouldMockUSeCaseCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"message":"invalid accountID received: 'invalid_account_id'. accountID must be not null, numeric and bigger than zero."}`,
	}
}

func makeScenarioThatReceivesAnErrorFromUseCaseLayerWhenTryToGetAnAccount() accountControllerGetScenario {
	return accountControllerGetScenario{
		TestName:              "Create an account",
		Method:                http.MethodGet,
		Uri:                   accountUri,
		PathParam:             "1",
		BodyString:            emptyBody,
		ShouldMockUSeCaseCall: true,
		MockUseCaseResult:     entities.Account{},
		MockUseCaseError:      errors.New("mock-error"),
		ExpectStatus:          http.StatusInternalServerError,
		ExpectResponse:        `{"message":"mock-error"}`,
	}
}

func makeScenarioThatCreateATransaction() transactionControllerCreateScenario {
	return transactionControllerCreateScenario{
		TestName:              "Create an account",
		Method:                http.MethodPost,
		Uri:                   transactionUri,
		BodyString:            `{"account_id":1,"operation_type_id":4,"amount":123.45}`,
		ShouldMockUSeCaseCall: true,
		MockUseCaseResult:     entities.Transaction{Model: gorm.Model{ID: 1}, AccountID: 1, OperationTypeID: 4, Amount: 123.45},
		MockUseCaseError:      nil,
		ExpectStatus:          http.StatusCreated,
		ExpectResponse:        `{}`,
	}
}

func makeScenarioThatReceivesABadRequestWhenTryToCreateATransactionSendingANonNumericAccountID() transactionControllerCreateScenario {
	return transactionControllerCreateScenario{
		TestName:              "Receives a bad request when send in 'account_id' a non-numeric value",
		Method:                http.MethodPost,
		Uri:                   transactionUri,
		BodyString:            `{"account_id":"1","operation_type_id":4,"amount":123.45}`,
		ShouldMockUSeCaseCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"message":"json: cannot unmarshal string into Go struct field CreateTransactionInput.account_id of type uint"}`,
	}
}

func makeScenarioThatReceivesABadRequestWhenTryToCreateATransactionSendingANonNumericOperationTypeID() transactionControllerCreateScenario {
	return transactionControllerCreateScenario{
		TestName:              "Receives a bad request when send in 'operation_type_id' a non-numeric value",
		Method:                http.MethodPost,
		Uri:                   transactionUri,
		BodyString:            `{"account_id":1,"operation_type_id":"4","amount":123.45}`,
		ShouldMockUSeCaseCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"message":"json: cannot unmarshal string into Go struct field CreateTransactionInput.operation_type_id of type uint"}`,
	}
}

func makeScenarioThatReceivesABadRequestWhenTryToCreateATransactionSendingANonNumericAmount() transactionControllerCreateScenario {
	return transactionControllerCreateScenario{
		TestName:              "Receives a bad request when send in 'amount' a non-numeric value",
		Method:                http.MethodPost,
		Uri:                   transactionUri,
		BodyString:            `{"account_id":1,"operation_type_id":4,"amount":"123.45"}`,
		ShouldMockUSeCaseCall: false,
		ExpectStatus:          http.StatusBadRequest,
		ExpectResponse:        `{"message":"json: cannot unmarshal string into Go struct field CreateTransactionInput.amount of type float64"}`,
	}
}

func makeScenarioThatReceivesAnErrorFromUseCaseLayerWhenTryToCreateATransaction() transactionControllerCreateScenario {
	return transactionControllerCreateScenario{
		TestName:              "Receives an error from use case layer",
		Method:                http.MethodPost,
		Uri:                   transactionUri,
		BodyString:            `{"account_id":1,"operation_type_id":4,"amount":123.45}`,
		ShouldMockUSeCaseCall: true,
		MockUseCaseResult:     entities.Transaction{},
		MockUseCaseError:      errors.New("mock-error"),
		ExpectStatus:          http.StatusInternalServerError,
		ExpectResponse:        `{"message":"mock-error"}`,
	}
}

func getFullUrl(uri string, pathParam string) (fullUrl string) {
	fullUrl = fmt.Sprintf("%s%s", baseUrl, uri)

	if pathParam != "" {
		fullUrl += fmt.Sprintf("/%s", pathParam)
	}

	return
}

func executeRequest(response *httptest.ResponseRecorder, method, requestUrl, body string, router *gin.Engine) {
	req, _ := http.NewRequest(method, requestUrl, bytes.NewBuffer([]byte(body)))
	router.ServeHTTP(response, req)
}
