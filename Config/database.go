package Config

import (
	"BearLibrary/Models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConnection() {
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "root")
	port := getEnv("DB_PORT", "3307")
	db := getEnv("DB_NAME", "library_of_bears")
	//dsn1 := "root:@tcp(host.docker.internal:3307)/library_of_bears?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, host, port, db)
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
	err = DB.AutoMigrate(&Models.User{})
	if err != nil {
		return
	}
}
