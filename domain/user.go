package domain

type User struct {
	BaseModel
	Email string `json:"email,omitempty"`
}
