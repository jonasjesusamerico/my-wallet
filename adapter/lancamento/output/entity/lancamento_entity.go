package entity

import (
	"time"

	"gorm.io/gorm"
)

type LancamentoEntity struct {
	ID             uint64    `gorm:"primaryKey"`
	DataCompra     time.Time `gorm:"type:date"`
	Descricao      string
	Setor          string
	FormaPagamento string
	Valor          float64
	Situacao       string
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
