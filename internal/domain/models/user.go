package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" validate:"required,min=3"`
	Username  string     `json:"username" gorm:"unique" validate:"required,min=3,max=30"`
	Email     string     `json:"email" gorm:"unique" validate:"required,email"`
	Password  string     `json:"password" validate:"required,min=8,max=32"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at,omitempty"`
	}{
		Alias:     (Alias)(u),
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: func() string {
			if u.DeletedAt != nil {
				return u.DeletedAt.Format("2006-01-02 15:04:05")
			}
			return ""
		}(),
	})
}
