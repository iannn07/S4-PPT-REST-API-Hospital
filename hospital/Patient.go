package hospital

type Patient struct {
	Patientid     int    `json:"PatientID"  gorm:"column:PatientID;primaryKey"`
	Doctorid      int    `json:"DoctorID"`
	Patientname   string `json:"PatientName"`
	Patientdob    string `json:"PatientDOB"`
	Patientgender string `json:"PatientGender"`
}
