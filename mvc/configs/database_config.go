package configs

import (
	"backend/mvc/models/doctors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:boomskii95.@tcp(127.0.0.1:3306)/rumah_sakit?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Databases Failed to connect")
	} else {
		fmt.Println("Database Connected")
	}
	migrationDB()
}

func migrationDB() {
	DB.AutoMigrate(&doctors.Doctors{})
}
