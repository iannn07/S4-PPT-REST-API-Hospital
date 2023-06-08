package hospital

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:5050)/hospital-management"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect!")
	}
	database = db
}
