package useCases

import (
	"digibank/internal/domain/entities"
	"digibank/internal/frameworks/errorx"
	"digibank/internal/interfaceAdapters/dto"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type accountUseCaseCreateScenario struct {
	TestName             string
	Data                 dto.CreateAccountInput
	MockRepoResult       entities.Account
	MockRepoError        error
	ExpectDocumentNumber string
	ExpectError          errorx.Errorx
}

type accountUseCaseGetScenario struct {
	TestName             string
	AccountID            uint
	MockRepoResult       entities.Account
	MockRepoError        error
	ExpectDocumentNumber string
	ExpectError          errorx.Errorx
}

type transactionUseCaseCreateScenario struct {
	TestName       string
	Data           dto.CreateTransactionInput
	MockRepoResult entities.Transaction
	MockRepoError  error
	ExpectResult   entities.Transaction
	ExpectError    errorx.Errorx
}

func makeScenarioThatCreateAnAccount() accountUseCaseCreateScenario {
	return accountUseCaseCreateScenario{
		TestName:             "Create an account",
		Data:                 dto.CreateAccountInput{DocumentNumber: "12345678900"},
		MockRepoResult:       entities.Account{DocumentNumber: "12345678900", Model: gorm.Model{ID: 1}},
		MockRepoError:        nil,
		ExpectDocumentNumber: "12345678900",
		ExpectError:          nil,
	}
}

func makeScenarioThatReceivesAnErrorWhenCreateAnAccount() accountUseCaseCreateScenario {
	return accountUseCaseCreateScenario{
		TestName:             "Create an account and receives an error from DB",
		Data:                 dto.CreateAccountInput{DocumentNumber: "12345678900"},
		MockRepoResult:       entities.Account{},
		MockRepoError:        errors.New("mock-error"),
		ExpectDocumentNumber: "",
		ExpectError:          errorx.NewErrorx(http.StatusInternalServerError, errors.New("mock-error")),
	}
}

func makeScenarioThatRetrievesAnAccount() accountUseCaseGetScenario {
	return accountUseCaseGetScenario{
		TestName:             "Get an account with accountID 1",
		AccountID:            1,
		MockRepoResult:       entities.Account{Model: gorm.Model{ID: 1}, DocumentNumber: "12345678900"},
		MockRepoError:        nil,
		ExpectDocumentNumber: "12345678900",
		ExpectError:          nil,
	}
}

func makeScenarioThatRetrievesNoneAccount() accountUseCaseGetScenario {
	return accountUseCaseGetScenario{
		TestName:             "Get an account with accountID 1 but receives none results from DB",
		AccountID:            1,
		MockRepoResult:       entities.Account{},
		MockRepoError:        errors.New("mock-error"),
		ExpectDocumentNumber: "",
		ExpectError:          errorx.NewErrorx(http.StatusInternalServerError, errors.New("mock-error")),
	}
}

func makeScenarioThatReceivesAnErrorWhenGetAnAccount() accountUseCaseGetScenario {
	return accountUseCaseGetScenario{
		TestName:             "Get an account with accountID 1 but receives an error from DB",
		AccountID:            1,
		MockRepoError:        errors.New("mock-error"),
		ExpectDocumentNumber: "",
		ExpectError:          errorx.NewErrorx(http.StatusInternalServerError, errors.New("mock-error")),
	}
}
func makeScenarioThatCreateATransaction() transactionUseCaseCreateScenario {
	return transactionUseCaseCreateScenario{
		TestName:       "Create an account",
		Data:           dto.CreateTransactionInput{AccountID: 1, OperationTypeID: 4, Amount: 123.45},
		MockRepoResult: entities.Transaction{Model: gorm.Model{ID: 1}, AccountID: 1, OperationTypeID: 4, Amount: 123.45},
		MockRepoError:  nil,
		ExpectResult:   entities.Transaction{AccountID: 1, OperationTypeID: 4, Amount: 123.45},
		ExpectError:    nil,
	}
}

func makeScenarioThatReceivesAnErrorWhenCreateATransaction() transactionUseCaseCreateScenario {
	return transactionUseCaseCreateScenario{
		TestName:      "Create an account and receives an error from DB",
		Data:          dto.CreateTransactionInput{AccountID: 1, OperationTypeID: 4, Amount: 123.45},
		MockRepoError: errors.New("mock-error"),
		ExpectResult:  entities.Transaction{},
		ExpectError:   errorx.NewErrorx(http.StatusInternalServerError, errors.New("mock-error")),
	}
}
