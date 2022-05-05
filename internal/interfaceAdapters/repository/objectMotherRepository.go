package repository

import (
	"digibank/internal/domain/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os/exec"
	"time"
)

const (
	testDbName = "test.db"
)

func beforeEach() *gorm.DB {
	removeTestDB()

	db, err := gorm.Open(sqlite.Open(testDbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("error to connect testDB. Err: %v", err)
	}

	_ = db.AutoMigrate(&entities.Account{})
	_ = db.AutoMigrate(&entities.Transaction{})
	return db
}

func removeTestDB() {
	cmd := exec.Command("rm", "-f", testDbName)
	if err := cmd.Run(); err != nil {
		log.Fatalf("error to remove previous DB. Err: %v", err)
	}
}

type accountRepositoryScenario struct {
	TestName             string
	Data                 entities.Account
	MockRows             []entities.Account
	ExpectID             uint
	ExpectDocumentNumber string
	ExpectError          error
}

type transactionRepositoryScenario struct {
	TestName    string
	Data        entities.Transaction
	MockRows    []entities.Transaction
	ExpectID    uint
	ExpectError error
}

func makeScenarioThatRetrievesAnAccount() accountRepositoryScenario {
	return accountRepositoryScenario{
		TestName:             "Get an account",
		Data:                 entities.Account{Model: gorm.Model{ID: 1}, DocumentNumber: "12345678900"},
		MockRows:             []entities.Account{{Model: gorm.Model{ID: 1}, DocumentNumber: "12345678900"}},
		ExpectID:             1,
		ExpectDocumentNumber: "12345678900",
	}
}

func makeScenarioThatCreatesAnAccount() accountRepositoryScenario {
	return accountRepositoryScenario{
		TestName: "Create an account",
		Data:     entities.Account{DocumentNumber: "12345678900"},
		ExpectID: 1,
	}
}

func makeScenarioThatCreatesATransaction() transactionRepositoryScenario {
	return transactionRepositoryScenario{
		TestName: "Create a transaction",
		Data: entities.Transaction{
			Model:           gorm.Model{ID: 1},
			AccountID:       1,
			OperationTypeID: 4,
			Amount:          123.45,
			EventDate:       time.Now(),
		},
		ExpectID: 1,
	}
}
