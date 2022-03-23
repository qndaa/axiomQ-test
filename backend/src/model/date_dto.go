package model

import (
	"encoding/json"
	"strings"
	"time"
)

type DateDTO time.Time

func (j *DateDTO) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		*j = DateDTO(time.Time{})
		return nil
	}
	*j = DateDTO(t)
	return nil
}

func (j DateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j DateDTO) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
