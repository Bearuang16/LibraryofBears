package Models

import "gorm.io/datatypes"
import "time"

type Arts struct {
	ID          uint           `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title       string         `json:"title"`
	Tags        datatypes.JSON `json:"tags"`
	AuthorID    uint           `json:"author_id"`
	File        string         `json:"file"`
	Source      string         `json:"source"`
	Website     string         `json:"website"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   time.Time      `json:"deleted_at"`
	SeriesID    uint           `json:"series_id"`
	Author      Author         `json:"author" gorm:"foreignkey:AuthorID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
	Series      Series         `json:"series" gorm:"foreignkey:SeriesID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
