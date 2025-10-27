package exttypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringArray []string

func (s *StringArray) FromDB(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, s)
}

func (s *StringArray) ToDB() ([]byte, error) {
	return json.Marshal(s)
}

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	if len(bytes) == 0 {
		*s = StringArray{}
		return nil
	}
	return json.Unmarshal(bytes, s)
}

func (s StringArray) Value() (driver.Value, error) {
	if len(s) == 0 {
		return json.Marshal([]string{})
	}
	return json.Marshal(s)
}

type Uint64Array []uint64

func (u *Uint64Array) FromDB(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, u)
}

func (u *Uint64Array) ToDB() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Uint64Array) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	if len(bytes) == 0 {
		*u = Uint64Array{}
		return nil
	}
	return json.Unmarshal(bytes, u)
}

func (u Uint64Array) Value() (driver.Value, error) {
	if len(u) == 0 {
		return json.Marshal([]uint64{})
	}
	return json.Marshal(u)
}

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
