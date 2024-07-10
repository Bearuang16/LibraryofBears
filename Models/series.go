package Models

type Series struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title    string `json:"title"`
	AuthorID uint   `json:"author_id"`
	Author   Author `json:"author" gorm:"foreignkey:AuthorID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
