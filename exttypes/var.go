package exttypes

import "encoding/json"

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
