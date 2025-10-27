package exttypes

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"time"
)

type JsonTime time.Time

// 用于iris的path,params,query参数序列化
func JsonTimeConverter(value string) reflect.Value {
	v, err := time.ParseInLocation("20060102150405", value, time.Local)
	if err != nil {
		return reflect.ValueOf(JsonTime{})
	}
	return reflect.ValueOf(v)
}

const dataTimeFormat = "2006-01-02 15:04:05"

func GetJsonTime(t time.Time) JsonTime {
	return JsonTime(t)
}

func (j *JsonTime) ToTime() time.Time {
	return time.Time(*j)
}

func (j JsonTime) MarshalJSON() ([]byte, error) {
	if time.Time(j).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(j).Format(dataTimeFormat) + `"`), nil
}

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	raw := strings.Trim(string(b), "\"")
	now, err := time.ParseInLocation(dataTimeFormat, raw, time.Local)
	if err != nil {
		j = &JsonTime{}
		return nil
	}
	*j = JsonTime(now)
	return nil
}

func (j *JsonTime) FromDB(b []byte) error {
	j.UnmarshalJSON(b)
	return nil
}

func (j *JsonTime) ToDB() ([]byte, error) {
	raw := time.Time(*j).Format(dataTimeFormat)
	if raw == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(raw), nil
}

// Value 实现 driver.Valuer 接口，用于数据库写入（兼容 PostgreSQL）
func (j JsonTime) Value() (driver.Value, error) {
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
func (j *JsonTime) Scan(value interface{}) error {
	if value == nil {
		*j = JsonTime{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*j = JsonTime(v)
		return nil
	case []byte:
		return j.FromDB(v)
	case string:
		return j.FromDB([]byte(v))
	default:
		*j = JsonTime{}
		return nil
	}
}

func (j JsonTime) IsZero() bool {
	return time.Time(j).IsZero()
}

func (j JsonTime) String() string {
	return time.Time(j).Format(dataTimeFormat)
}

func (j JsonTime) Format(layout string) string {
	return time.Time(j).Format(layout)
}

func (j JsonTime) ToJsonDate() JsonDate {
	return JsonDate(time.Time(j))
}

func Now() JsonTime {
	return JsonTime(time.Now())
}
