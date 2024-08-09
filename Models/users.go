package Models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
