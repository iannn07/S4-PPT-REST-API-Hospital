package Models

type User struct {
	ID       int    `gorm:"column:id_user;primaryKey;autoIncrement"  json:"id_user"`
	Username string `gorm:"column:username;type:varchar(255);unique" json:"username"`
	Password string `gorm:"column:password;type:varchar(255)"        json:"password"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}