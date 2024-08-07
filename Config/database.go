package Config

import (
	"BearLibrary/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConnection() {
	dsn := "root:@tcp(host.docker.internal:3307)/library_of_bears?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	err := DB.AutoMigrate(&Models.Author{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&Models.Series{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&Models.Arts{})
	if err != nil {
		return
	}
}
