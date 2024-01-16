package converter

import (
	"my-wallet/adapter/lancamento/input/model/response"
	"my-wallet/application/domain"
)

func ConvertDomainToResponse(
	lancamentoDomain *domain.LancamentoDomain,
) response.LancamentoResponse {
	return response.LancamentoResponse{
		ID:             lancamentoDomain.ID,
		DataCompra:     lancamentoDomain.DataCompra,
		Descricao:      lancamentoDomain.Descricao,
		Setor:          lancamentoDomain.Setor,
		FormaPagamento: lancamentoDomain.FormaPagamento,
		Valor:          lancamentoDomain.Valor,
		Situacao:       lancamentoDomain.Situacao,
	}
}
