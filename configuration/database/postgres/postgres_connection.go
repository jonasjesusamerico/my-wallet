package postgres

import (
	"context"
	"log"
	"my-wallet/adapter/lancamento/output/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

type Database struct {
}

func NewPostgresDBConnection(ctx context.Context) (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=my-wallet port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}
	DB.AutoMigrate(&entity.LancamentoEntity{})

	return DB, nil
}
