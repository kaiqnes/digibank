package dependencyInjection

import (
	"digibank/internal/interfaceAdapters/controllers"
	"digibank/internal/interfaceAdapters/presenters"
	"digibank/internal/interfaceAdapters/repository"
	"digibank/internal/useCases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dependencyInjection struct {
	routes *gin.Engine
	db     *gorm.DB
}

func NewDependencyInjection(routes *gin.Engine, session *gorm.DB) *dependencyInjection {
	return &dependencyInjection{
		routes: routes,
		db:     session,
	}
}

func (di *dependencyInjection) SetupDependencies() {
	di.injectStructuralResources()
	di.injectPublicResources()
}

func (di *dependencyInjection) injectPublicResources() {
	publicGroup := di.routes.Group("/api/v1")

	errPresenter := presenters.NewErrorPresenter()

	/* Accounts */
	accountsPresenter := presenters.NewAccountPresenter(errPresenter)
	accountsRepository := repository.NewAccountsRepository(di.db)
	accountsUseCase := useCases.NewAccountsUseCase(accountsRepository)
	accounts := controllers.NewAccountsController(publicGroup, accountsPresenter, accountsUseCase)
	accounts.SetupEndpoints()

	/* Transactions */
	transactionsPresenter := presenters.NewTransactionPresenter(errPresenter)
	transactionsRepository := repository.NewTransactionsRepository(di.db)
	transactionsUseCase := useCases.NewTransactionUseCase(transactionsRepository)
	transactions := controllers.NewTransactionsController(publicGroup, transactionsPresenter, transactionsUseCase)
	transactions.SetupEndpoints()
}

func (di *dependencyInjection) injectStructuralResources() {
	healthCheck := controllers.NewHealthCheckController(di.routes)
	healthCheck.SetupEndpoints()

	swagger := controllers.NewSwaggerController(di.routes)
	swagger.SetupEndpoints()
}
