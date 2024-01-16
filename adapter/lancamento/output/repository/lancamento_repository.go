package repository

import (
	"my-wallet/adapter/lancamento/output/converter"
	"my-wallet/adapter/lancamento/output/entity"
	"my-wallet/application/domain"
	"my-wallet/application/port/lancamento/output"
	logger "my-wallet/configuration"
	"my-wallet/configuration/rest_errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewLancamentoRepository(database *gorm.DB) output.LancamentoPort {
	return &lancamentoRepository{
		database,
	}
}

type lancamentoRepository struct {
	db *gorm.DB
}

// FindLancamentoByID implements output.LancamentoPort.
func (lc *lancamentoRepository) FindLancamentoByID(id uint64) (*domain.LancamentoDomain, *rest_errors.RestErr) {
	logger.Info("Init createLancamento repository", zap.String("journey", "findLancamentoById"))
	lancamentoEntity := &entity.LancamentoEntity{}
	if err := lc.db.First(&lancamentoEntity, id).Error; err != nil {
		return nil, rest_errors.NewNotFoundError(err.Error())
	}

	return converter.ConvertEntityToDomain(*lancamentoEntity), nil
}

func (ur *lancamentoRepository) CreateLancamento(lancamentoDomain domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr) {
	logger.Info("Init createLancamento repository", zap.String("journey", "createLancamento"))

	value := converter.ConvertDomainToEntity(lancamentoDomain)

	ur.db.Create(value)

	logger.Info(
		"CreateLancamento repository executed successfully",
		zap.String("lancamentoId", string(rune(value.ID))),
		zap.String("journey", "createLancamento"))

	return converter.ConvertEntityToDomain(*value), nil
}
