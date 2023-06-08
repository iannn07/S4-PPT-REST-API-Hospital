package hospital

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/finpro_ppt"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect!")
	}
	DB = db
}
