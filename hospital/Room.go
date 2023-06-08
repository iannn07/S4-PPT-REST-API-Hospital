package hospital

type Room struct {
	roomid    int    `json:"RoomID"  gorm:"primary_key"`
	patientid int    `json:"PatientID"`
	roomtype  string `json:"RoomType"`
}
