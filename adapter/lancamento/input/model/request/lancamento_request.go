package request

import (
	"my-wallet/configuration/custom/date"
)

type LancamentoRequest struct {
	DataCompra     date.CustomTime `json:"dataCompra"`
	Descricao      string          `json:"descricao"`
	Setor          string          `json:"setor"`
	FormaPagamento string          `json:"formaPagamento"`
	Valor          float64         `json:"valor"`
	Situacao       string          `json:"situacao"`
}
