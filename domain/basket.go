package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Basket struct {
	BaseModel
	Data   Data        `json:"data"`
	State  BasketState `json:"state"`
	UserID uint        `json:"user_id"`
}
type BasketState string

const (
	BasketStatePending  BasketState = "PENDING"
	BasketStateComplete BasketState = "COMPLETED"
)

type Data map[string]interface{}

func (d Data) Value() (driver.Value, error) {
	valueStr, err := json.Marshal(d)
	return string(valueStr), err
}

func (d *Data) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &d)
}
