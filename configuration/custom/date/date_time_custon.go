package date

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type CustomTime struct {
	time.Time
}

const customTimeLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	if s == "null" {
		return nil
	}

	t, err := time.Parse(`"`+customTimeLayout+`"`, s)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

func (ct *CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	ct.Time = value.(time.Time)
	return nil
}

func (d *CustomTime) BeforeSave(tx *gorm.DB) error {
	// Formatar o campo DataCompra antes de salvar no banco de dados
	d.Time = d.UTC()
	return nil
}
