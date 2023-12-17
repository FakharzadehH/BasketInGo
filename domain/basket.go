package domain

type Basket struct {
	BaseModel
	Data   string      `json:"data,omitempty"`
	State  BasketState `json:"state,omitempty"`
	UserID uint        `json:"user_id,omitempty"`
}
type BasketState string

const (
	BasketStatePending  BasketState = "PENDING"
	BasketStateComplete BasketState = "COMPLETED"
)
