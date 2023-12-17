package domain

type User struct {
	BaseModel
	Username string `json:"email,omitempty"`
}
