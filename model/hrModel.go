package model

import (
	"time"
)

type Hr struct {
	ID        uint      `json:"id" gorm:"primaryKey"`                      // Primary key
	Name      string    `json:"name" gorm:"not null;column:name;size:100"` // Name column with constraints
	Email     string    `json:"email" gorm:"not null;column:email;size:100"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically set when the record is created
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
