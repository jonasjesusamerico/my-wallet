package controller

import (
	"fmt"
	"my-wallet/adapter/lancamento/input/converter"
	"my-wallet/adapter/lancamento/input/model/request"
	"my-wallet/application/domain"
	"my-wallet/application/port/lancamento/input"
	logger "my-wallet/configuration"
	"my-wallet/configuration/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LancamentoControllerInterface interface {
	CreateLancamento(c *gin.Context)
	FindLancamentoByID(c *gin.Context)
}

type lancamentoControllerInterface struct {
	service input.LancamentoDomainService
}

func NewLancamentoControllerInterface(service input.LancamentoDomainService) LancamentoControllerInterface {
	return &lancamentoControllerInterface{
		service: service,
	}
}

func (uc *lancamentoControllerInterface) CreateLancamento(c *gin.Context) {
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

	domainResult, err := uc.service.CreateLancamentoServices(lancamentoDomain)
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

func (lc *lancamentoControllerInterface) FindLancamentoByID(c *gin.Context) {
	logger.Info("Init findLancamentoByID controller",
		zap.String("journey", "findLancamentoByID"),
	)

	lancamentoId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		logger.Error(
			"Error when trying to parse lancamentoId for uint", err,
			zap.String("journey", "createLancamento"))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(lancamentoId)

	LancamentoDomain, errr := lc.service.FindByIdLancamentoServices(lancamentoId)
	if errr != nil {
		logger.Error("Error trying to call findLancamentoByID services",
			err,
			zap.String("journey", "findLancamentoByID"),
		)
		c.JSON(errr.Code, errr)
		return
	}

	logger.Info("FindLancamentoByID controller executed successfully",
		zap.String("journey", "findLancamentoByID"),
	)
	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(LancamentoDomain))
}
