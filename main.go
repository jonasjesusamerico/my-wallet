package main

import (
	"context"
	"log"
	"my-wallet/adapter/lancamento/input/controller"
	routeLancamento "my-wallet/adapter/lancamento/input/controller/routes"
	"my-wallet/adapter/lancamento/output/repository"
	service "my-wallet/application/services"
	logger "my-wallet/configuration"
	postgres "my-wallet/configuration/database/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	logger.Info("About to start lancamento application")

	database, _ := postgres.NewPostgresDBConnection(context.Background())

	ginDefault := gin.Default()
	router := ginDefault.Group("/api")

	{
		routeV1 := router.Group("/v1")
		routeLancamento.InitRoutesLancamento(routeV1, initLancamento(database))
	}

	if err := ginDefault.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initLancamento(database *gorm.DB) controller.LancamentoControllerInterface {
	lancamentoRepo := repository.NewLancamentoRepository(database)
	lancamentoService := service.NewLancamentoDomainService(lancamentoRepo)
	return controller.NewLancamentoControllerInterface(lancamentoService)
}
