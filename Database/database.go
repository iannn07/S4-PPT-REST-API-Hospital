package Database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"HospitalFinpro/Models"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/hospitals123"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect!")
	}
	
	db.AutoMigrate(
		&Models.Doctor{},
		&Models.User{},
		&Models.Patient{},
		&Models.Room{},
		&Models.Diagnose{},
		&Models.Payment{},
	)

	DB = db
}