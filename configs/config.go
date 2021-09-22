package configs

import (
	"backend/internal/daftar_kunjungan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "root:boomskii95.@tcp(127.0.0.1:3306)/rumah_sakit?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Databases Failed to connect")
	}
	Migrate(DB)
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.Patient{})
}
