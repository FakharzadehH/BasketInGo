package domain

type Basket struct {
	BaseModel
	Data string `json:"data,omitempty"`
}
