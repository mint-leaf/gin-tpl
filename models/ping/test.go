package ping

import "encoding/json"

type TestModel struct {
	ID uint64 `json:"id" xorm:"id"`
}

// implement for xorm
func (t *TestModel) TableName() string {
	return "test"
}

// implement for print, easy debug and log
func (t *TestModel) String() string {
	d, _ := json.Marshal(t)
	return string(d)
}
