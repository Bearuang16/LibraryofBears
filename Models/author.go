package Models

import "time"

type Author struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
