package domain

type Basket struct {
	BaseModel
	Data   string `json:"data,omitempty"`
	UserID uint   `json:"user_id,omitempty"`
}
