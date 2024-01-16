package converter

import (
	"my-wallet/adapter/lancamento/output/entity"
	"my-wallet/application/domain"
	"my-wallet/configuration/custom/date"
)

func ConvertEntityToDomain(entity entity.LancamentoEntity) *domain.LancamentoDomain {
	return &domain.LancamentoDomain{
		ID:             entity.ID,
		DataCompra:     date.CustomTime{Time: entity.DataCompra},
		Descricao:      entity.Descricao,
		Setor:          entity.Setor,
		FormaPagamento: entity.FormaPagamento,
		Valor:          entity.Valor,
		Situacao:       entity.Situacao,
	}
}
