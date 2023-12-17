package domain

type Basket struct {
	BaseModel
	Data   string      `json:"data"`
	State  BasketState `json:"state"`
	UserID uint        `json:"user_id"`
}
type BasketState string

const (
	BasketStatePending  BasketState = "PENDING"
	BasketStateComplete BasketState = "COMPLETED"
)
