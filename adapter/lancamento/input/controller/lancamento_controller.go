package controller

import (
	"my-wallet/adapter/lancamento/input/converter"
	"my-wallet/adapter/lancamento/input/model/request"
	"my-wallet/application/domain"
	"my-wallet/application/port/lancamento/input"
	logger "my-wallet/configuration"
	"my-wallet/configuration/rest_errors"
	"my-wallet/configuration/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LancamentoControllerInterface interface {
	Save(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
}

type lancamentoControllerInterface struct {
	service input.LancamentoDomainService
}

func NewLancamentoControllerInterface(service input.LancamentoDomainService) LancamentoControllerInterface {
	return &lancamentoControllerInterface{
		service: service,
	}
}

func (lc *lancamentoControllerInterface) Save(c *gin.Context) {
	logger.Info("Init CreateLancamento controller",
		zap.String("journey", "createLancamento"),
	)
	lancamentoRequest := request.LancamentoRequest{}

	if err := c.ShouldBindJSON(&lancamentoRequest); err != nil {
		logger.Error("Error trying to validate lancamento info", err,
			zap.String("journey", "createLancamento"))
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	lancamentoDomain := domain.LancamentoDomain{
		DataCompra:     lancamentoRequest.DataCompra,
		Descricao:      lancamentoRequest.Descricao,
		Setor:          lancamentoRequest.Setor,
		FormaPagamento: lancamentoRequest.FormaPagamento,
		Valor:          lancamentoRequest.Valor,
		Situacao:       lancamentoRequest.Situacao,
	}

	domainResult, err := lc.service.Save(lancamentoDomain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateLancamento service",
			err,
			zap.String("journey", "createLancamento"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateLancamento controller executed successfully",
		zap.String("lancamentoId", string(rune(domainResult.ID))),
		zap.String("journey", "createLancamento"))

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(
		domainResult,
	))
}

// Update implements LancamentoControllerInterface.
func (lc *lancamentoControllerInterface) Update(c *gin.Context) {
	logger.Info("Init UpdateLancamento controller",
		zap.String("journey", "updateLancamento"),
	)
	lancamentoRequest := request.LancamentoRequest{}

	if err := c.ShouldBindJSON(&lancamentoRequest); err != nil {
		logger.Error("Error trying to validate lancamento info", err,
			zap.String("journey", "createLancamento"))
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	lancamentoDomain := domain.LancamentoDomain{
		ID:             lancamentoRequest.ID,
		DataCompra:     lancamentoRequest.DataCompra,
		Descricao:      lancamentoRequest.Descricao,
		Setor:          lancamentoRequest.Setor,
		FormaPagamento: lancamentoRequest.FormaPagamento,
		Valor:          lancamentoRequest.Valor,
		Situacao:       lancamentoRequest.Situacao,
	}

	domainResult, restError := lc.service.Update(lancamentoDomain)

	if restError != nil {
		logger.Error(
			"Error trying to call UpdateLancamento service",
			restError,
			zap.String("journey", "createLancamento"))
		c.JSON(restError.Code, restError)
		return
	}

	logger.Info(
		"UpdateLancamento controller executed successfully",
		zap.String("lancamentoId", string(rune(domainResult.ID))),
		zap.String("journey", "createLancamento"))

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(
		domainResult,
	))

}

func (lc *lancamentoControllerInterface) FindById(c *gin.Context) {
	logger.Info("Init findLancamentoByID controller",
		zap.String("journey", "findLancamentoByID"),
	)

	lancamentoId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		logger.Error(
			"Error when trying to parse lancamentoId for uint", err,
			zap.String("journey", "createLancamento"))

		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError(err.Error()))
		return
	}

	LancamentoDomain, errRest := lc.service.FindById(lancamentoId)
	if errRest != nil {
		logger.Error("Error trying to call findLancamentoByID services",
			err,
			zap.String("journey", "findLancamentoByID"),
		)
		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("FindLancamentoByID controller executed successfully",
		zap.String("journey", "findLancamentoByID"),
	)
	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(LancamentoDomain))
}
