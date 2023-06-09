package hospital

type Room struct {
	RoomID    int    `json:"RoomID"  gorm:"column:RoomID;primaryKey"`
	PatientID int    `json:"PatientID"`
	RoomType  string `json:"RoomType"`
}
