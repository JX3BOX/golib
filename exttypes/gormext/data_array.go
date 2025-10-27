package gormext

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DataArray[T any] []T

func (u *DataArray[T]) FromDB(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, u)
}

func (u *DataArray[T]) ToDB() ([]byte, error) {
	return json.Marshal(u)
}

func (u *DataArray[T]) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	if len(bytes) == 0 {
		*u = DataArray[T]{}
		return nil
	}
	return json.Unmarshal(bytes, u)
}

func (u DataArray[T]) Value() (driver.Value, error) {
	if len(u) == 0 {
		return json.Marshal([]T{})
	}
	return json.Marshal(u)
}

// GormDataType gorm common data type
func (DataArray[T]) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (DataArray[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
