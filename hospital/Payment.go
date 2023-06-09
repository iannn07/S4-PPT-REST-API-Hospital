package hospital

type Payment struct {
	PayID     int    `json:"PayID"  gorm:"column:PayID;primaryKey"`
	PatientID int    `json:"PatientID"`
	PayTotal  string `json:"PayTotal"`
}
