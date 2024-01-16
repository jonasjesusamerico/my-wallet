package routes

import (
	"my-wallet/adapter/lancamento/input/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutesLancamento(r *gin.RouterGroup, lancamentoController controller.LancamentoControllerInterface) {
	lancamentoGroup := r.Group("/lancamentos")

	lancamentoGroup.POST("/", lancamentoController.Save)
	lancamentoGroup.GET("/:id", lancamentoController.FindById)
}
