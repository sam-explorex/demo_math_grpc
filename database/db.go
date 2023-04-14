package database

import (
	"github.com/sam-explorex/demo_math_grpc/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// connect db
	dsn := "root:password@tcp(127.0.0.1:57472)/demo_math_db?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return Db
}

func MigrateDatabase(Db *gorm.DB) {
	// migrate tables
	err := Db.AutoMigrate(&entities.Logs{})
	if err != nil {
		panic("Failed to auto-migrate model!")
	}
}
