package domain

import "time"

type BaseModel struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}