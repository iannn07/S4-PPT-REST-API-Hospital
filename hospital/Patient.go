package hospital

type Patient struct {
	PatientID     int    `json:"PatientID"  gorm:"column:PatientID;primaryKey"`
	DoctorID      int    `json:"DoctorID"`
	PatientName   string `json:"PatientName"`
	PatientDOB    string `json:"PatientDOB"`
	PatientGender string `json:"PatientGender"`
}
