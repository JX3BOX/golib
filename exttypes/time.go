package exttypes

import (
	"strings"
	"time"
)

type JsonTime time.Time

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

func (j JsonTime) IsZero() bool {
	return time.Time(j).IsZero()
}

func (j JsonTime) String() string {
	return time.Time(j).Format(dataTimeFormat)
}

func (j JsonTime) Format(layout string) string {
	return time.Time(j).Format(layout)
}

func Now() JsonTime {
	return JsonTime(time.Now())
}
