package domain

type Basket struct {
	BaseModel
	Data   string      `json:"data,omitempty"`
	State  basketState `json:"state,omitempty"`
	UserID uint        `json:"user_id,omitempty"`
}
type basketState string

const (
	basketStatePending  basketState = "PENDING"
	basketStateComplete basketState = "COMPLETED"
)
