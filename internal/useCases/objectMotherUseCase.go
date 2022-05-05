package useCases

import (
	"digibank/internal/domain/entities"
	"digibank/internal/interfaceAdapters/dto"
	"errors"
	"gorm.io/gorm"
)

type accountUseCaseCreateScenario struct {
	TestName             string
	Data                 dto.CreateAccountInput
	MockRepoError        error
	ExpectDocumentNumber string
	ExpectError          error
}

type accountUseCaseGetScenario struct {
	TestName             string
	AccountID            uint
	MockRepoResult       entities.Account
	MockRepoError        error
	ExpectDocumentNumber string
	ExpectError          error
}

func makeScenarioThatCreateAnAccount() accountUseCaseCreateScenario {
	return accountUseCaseCreateScenario{
		TestName:             "Create an account",
		Data:                 dto.CreateAccountInput{DocumentNumber: "12345678900"},
		MockRepoError:        nil,
		ExpectDocumentNumber: "12345678900",
		ExpectError:          nil,
	}
}

func makeScenarioThatReceivesAnErrorWhenCreateAnAccount() accountUseCaseCreateScenario {
	return accountUseCaseCreateScenario{
		TestName:             "Create an account and receives an error from DB",
		Data:                 dto.CreateAccountInput{DocumentNumber: "12345678900"},
		MockRepoError:        errors.New("mock-error"),
		ExpectDocumentNumber: "12345678900",
		ExpectError:          errors.New("mock-error"),
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
		ExpectError:          errors.New("mock-error"),
	}
}

func makeScenarioThatReceivesAnErrorWhenGetAnAccount() accountUseCaseGetScenario {
	return accountUseCaseGetScenario{
		TestName:             "Get an account with accountID 1 but receives an error from DB",
		AccountID:            1,
		MockRepoError:        errors.New("mock-error"),
		ExpectDocumentNumber: "",
		ExpectError:          errors.New("mock-error"),
	}
}
