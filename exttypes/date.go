package exttypes

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"time"
)

type JsonDate time.Time

// 用于iris的path,params,query参数序列化
func JsonDateConverter(value string) reflect.Value {
	v, err := time.ParseInLocation("20060102", value, time.Local)
	if err != nil {
		return reflect.ValueOf(JsonDate{})
	}
	return reflect.ValueOf(v)
}

const dateFormat = "2006-01-02"

func GetJsonDate(t time.Time) JsonDate {
	return JsonDate(t)
}

func (j *JsonDate) ToTime() time.Time {
	return time.Time(*j)
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	if time.Time(j).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(j).Format(dateFormat) + `"`), nil
}

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	raw := strings.Trim(string(b), "\"")
	now, err := time.ParseInLocation(dateFormat, raw, time.Local)
	if err != nil {
		j = &JsonDate{}
		return nil
	}
	*j = JsonDate(now)
	return nil
}

func (j *JsonDate) FromDB(b []byte) error {
	j.UnmarshalJSON(b)
	return nil

}

func (j *JsonDate) ToDB() ([]byte, error) {
	raw := time.Time(*j).Format(dateFormat)
	if raw == "0001-01-01" {
		return nil, nil
	}
	return []byte(raw), nil
}

// Value 实现 driver.Valuer 接口，用于数据库写入（兼容 PostgreSQL）
func (j JsonDate) Value() (driver.Value, error) {
	if j.IsZero() {
		return nil, nil
	}
	t := time.Time(j)
	if t.Year() == 1 {
		return nil, nil
	}
	return t, nil
}

// Scan 实现 sql.Scanner 接口，用于数据库读取（兼容 PostgreSQL）
func (j *JsonDate) Scan(value interface{}) error {
	if value == nil {
		*j = JsonDate{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*j = JsonDate(v)
		return nil
	case []byte:
		return j.FromDB(v)
	case string:
		return j.FromDB([]byte(v))
	default:
		*j = JsonDate{}
		return nil
	}
}

func (j JsonDate) IsZero() bool {
	return time.Time(j).IsZero()
}

func (j JsonDate) String() string {
	return time.Time(j).Format(dateFormat)
}

func (j JsonDate) Format(layout string) string {
	return time.Time(j).Format(layout)
}

func (j JsonDate) ToJsonTime() JsonTime {
	return JsonTime(j.ToTime())
}
