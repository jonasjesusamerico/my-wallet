package service

import (
	"my-wallet/application/domain"
	"my-wallet/application/port/lancamento/input"
	"my-wallet/application/port/lancamento/output"
	logger "my-wallet/configuration"
	"my-wallet/configuration/rest_errors"

	"go.uber.org/zap"
)

func NewLancamentoDomainService(
	lancamentoRepository output.LancamentoPort) input.LancamentoDomainService {
	return &lancamentoDomainService{
		lancamentoRepository,
	}
}

type lancamentoDomainService struct {
	repository output.LancamentoPort
}

func (ud *lancamentoDomainService) CreateLancamentoServices(lancamentoDomain domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr) {

	logger.Info("Init createLancamento model.", zap.String("journey", "createLancamento"))

	lancamentoDomainRepository, err := ud.repository.CreateLancamento(lancamentoDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createLancamento"))
		return nil, err
	}

	logger.Info(
		"CreateLancamento service executed successfully",
		zap.Uint("lancamentoId", lancamentoDomainRepository.ID),
		zap.String("journey", "createLancamento"))
	return lancamentoDomainRepository, nil
}

func (ud *lancamentoDomainService) FindByIdLancamentoServices(id uint64) (lancamento *domain.LancamentoDomain, err *rest_errors.RestErr) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	lancamento, err = ud.repository.FindLancamentoByID(id)
	return
}
