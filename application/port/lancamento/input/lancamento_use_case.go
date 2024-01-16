package input

import (
	"my-wallet/application/domain"
	"my-wallet/configuration/rest_errors"
)

type LancamentoDomainService interface {
	Save(domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr)
	Update(domain.LancamentoDomain) (*domain.LancamentoDomain, *rest_errors.RestErr)
	FindById(uint64) (*domain.LancamentoDomain, *rest_errors.RestErr)
}
