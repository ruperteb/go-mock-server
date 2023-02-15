package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *gorm.DB

func InitDB() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB = db

	DB.AutoMigrate(&User{})

	fmt.Println("test1", DB.Config)

}
