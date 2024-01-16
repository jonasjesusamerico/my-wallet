package output

import (
	"my-wallet/application/domain"
	"my-wallet/configuration/rest_errors"
)

type LancamentoPort interface {
	CreateLancamento(userDomain domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr)
	FindLancamentoByID(id uint64) (*domain.LancamentoDomain, *rest_errors.RestErr)
}
