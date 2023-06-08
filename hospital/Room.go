package hospital

type Room struct {
	Roomid    int    `json:"RoomID"  gorm:"column:RoomID;primaryKey"`
	Patientid int    `json:"PatientID"`
	Roomtype  string `json:"RoomType"`
}
