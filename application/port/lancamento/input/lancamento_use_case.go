package input

import (
	"my-wallet/application/domain"
	"my-wallet/configuration/rest_errors"
)

type LancamentoDomainService interface {
	CreateLancamentoServices(domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr)
	FindByIdLancamentoServices(uint64) (*domain.LancamentoDomain, *rest_errors.RestErr)
}
