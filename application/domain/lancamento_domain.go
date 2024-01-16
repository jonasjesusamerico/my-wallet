package domain

import (
	"my-wallet/configuration/custom/date"
)

type LancamentoDomain struct {
	ID             uint            `json:"id"`
	DataCompra     date.CustomTime `json:"dataCompra"`
	Descricao      string          `json:"descricao"`
	Setor          string          `json:"setor"`
	FormaPagamento string          `json:"formaPagamento"`
	Valor          float64         `json:"valor"`
	Situacao       string          `json:"situacao"`
}
