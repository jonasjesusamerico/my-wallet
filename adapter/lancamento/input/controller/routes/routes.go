package routes

import (
	"my-wallet/adapter/lancamento/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutesLancamento(r *gin.RouterGroup, lancamentoController controller.LancamentoControllerInterface) {
	lancamentoGroup := r.Group("/lancamentos")

	lancamentoGroup.POST("/", lancamentoController.CreateLancamento)
	lancamentoGroup.GET("/:id", lancamentoController.FindLancamentoByID)
}
