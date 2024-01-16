package output

import (
	"my-wallet/application/domain"
	"my-wallet/configuration/rest_errors"
)

type LancamentoPort interface {
	Save(userDomain domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr)
	FindById(id uint64) (*domain.LancamentoDomain, *rest_errors.RestErr)
}
