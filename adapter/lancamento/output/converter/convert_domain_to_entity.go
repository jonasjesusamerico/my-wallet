package converter

import (
	"my-wallet/adapter/lancamento/output/entity"
	"my-wallet/application/domain"
)

func ConvertDomainToEntity(domain domain.LancamentoDomain) *entity.LancamentoEntity {
	return &entity.LancamentoEntity{
		ID:             domain.ID,
		DataCompra:     domain.DataCompra.Time,
		Descricao:      domain.Descricao,
		Setor:          domain.Setor,
		FormaPagamento: domain.FormaPagamento,
		Valor:          domain.Valor,
		Situacao:       domain.Situacao,
	}
}
