package dependencyInjection

import (
	"digibank/internal/interfaceAdapters/controllers"
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

	/* Accounts */
	//accountsPresenter := presenters.NewAccountsPresenter()
	//accountsRepository := repository.NewAccountsRepository()
	//accountsUseCase := useCases.NewAccountsUseCase()
	accounts := controllers.NewAccountsController(publicGroup, di.db)
	accounts.SetupEndpoints()

	/* Transactions */
	//transactionsPresenter := presenters.NewTransactionsPresenter()
	//transactionsRepository := repository.NewTransactionsRepository()
	//transactionsUseCase := useCases.NewTransactionsUseCase()
	transactions := controllers.NewTransactionsController(publicGroup, di.db)
	transactions.SetupEndpoints()
}

func (di *dependencyInjection) injectStructuralResources() {
	healthCheck := controllers.NewHealthCheckController(di.routes)
	healthCheck.SetupEndpoints()

	swagger := controllers.NewSwaggerController(di.routes)
	swagger.SetupEndpoints()
}
