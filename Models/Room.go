package Models

type Room struct {
	ID 		    int     `gorm:"column:room_id;primaryKey;autoIncrement"			json:"room_id"`
	PatientID 	int     `gorm:"column:patient_id"								json:"patient_id"`
	RoomType  	string  `gorm:"column:room_type;varchar(255)"					json:"room_type"`
	Patient   	Patient `gorm:"foreignKey:PatientID"							json:"patient"`
}

type RoomInput struct {
	PatientID int    `json:"patient_id"`
	RoomType  string `json:"room_type"`
}

type RoomResponse struct {
	ID 		    int    `json:"room_id"`
	PatientID 	int    `json:"patient_id"`
	RoomType  	string `json:"room_type"`
}

type RoomListResponse struct {
	Rooms []RoomResponse `json:"rooms"`
}
