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

func (ls *lancamentoDomainService) Save(lancamentoDomain domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr) {

	logger.Info("Init createLancamento model.", zap.String("journey", "createLancamento"))

	lancamentoDomainRepository, err := ls.repository.Save(lancamentoDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createLancamento"))
		return nil, err
	}

	logger.Info(
		"CreateLancamento service executed successfully",
		zap.Uint64("lancamentoId", lancamentoDomainRepository.ID),
		zap.String("journey", "createLancamento"))
	return lancamentoDomainRepository, nil
}

func (ls *lancamentoDomainService) Update(lancamentoDomain domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr) {

	logger.Info("Init updateLancamento model.", zap.String("journey", "updateLancamento"))

	lancamentoDomainRepository, err := ls.repository.Update(lancamentoDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateLancamento"))
		return nil, err
	}

	logger.Info(
		"UpdateLancamento service executed successfully",
		zap.Uint64("lancamentoId", lancamentoDomainRepository.ID),
		zap.String("journey", "updateLancamento"))
	return lancamentoDomainRepository, nil
}

func (ls *lancamentoDomainService) FindById(id uint64) (lancamento *domain.LancamentoDomain, err *rest_errors.RestErr) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	lancamento, err = ls.repository.FindById(id)
	return
}
