package Models

type Author struct {
	ID   uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}
