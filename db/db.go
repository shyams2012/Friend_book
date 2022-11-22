package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql database connection
func DbConn() *gorm.DB {
	dbUserName := os.Getenv("MYSQL_USERNAME")
	dbPassword := os.Getenv("MYSQL_PASS")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/friend_book?charset=utf8mb4&parseTime=True&loc=Local", dbUserName, dbPassword)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
