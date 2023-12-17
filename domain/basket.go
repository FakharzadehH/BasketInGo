package domain

type Basket struct {
	BaseModel
	Data   string      `json:"data,omitempty"`
	State  BasketState `json:"state,omitempty"`
	UserID uint        `json:"user_id,omitempty"`
}
type BasketState string

const (
	basketStatePending  BasketState = "PENDING"
	basketStateComplete BasketState = "COMPLETED"
)
