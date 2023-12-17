package domain

type User struct {
	BaseModel
	Username string `json:"username"`
}
