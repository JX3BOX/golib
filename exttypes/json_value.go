package exttypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JsonValue struct {
	JsonByte []byte
	JsonData interface{}
	Content  []byte
}

func (j *JsonValue) FromDB(b []byte) error {
	j.JsonByte = make([]byte, len(b))
	copy(j.JsonByte, b)
	return json.Unmarshal(b, &j.JsonData)
}

func (j *JsonValue) ToDB() ([]byte, error) {
	if j.Content != nil && len(j.Content) > 0 {
		return j.Content, nil
	}
	return json.Marshal(j.JsonData)
}

func (j *JsonValue) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	if len(bytes) == 0 {
		j.JsonData = nil
		return nil
	}
	j.JsonByte = make([]byte, len(bytes))
	copy(j.JsonByte, bytes)
	return json.Unmarshal(bytes, &j.JsonData)
}

func (j JsonValue) Value() (driver.Value, error) {
	if j.JsonData == nil {
		return nil, nil
	}
	if j.Content != nil && len(j.Content) > 0 {
		return j.Content, nil
	}
	return json.Marshal(j.JsonData)
}

func (j JsonValue) MarshalJSON() ([]byte, error) {
	if len(j.JsonByte) == 0 {
		return json.Marshal(j.JsonData)
	}
	return j.JsonByte, nil
}

func (j *JsonValue) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &j.JsonData)
}

func (j *JsonValue) IsEmpty() bool {
	d, e := j.MarshalJSON()
	if e != nil {
		return true
	}
	if string(d) == "null" {
		return true
	}
	return false
}
