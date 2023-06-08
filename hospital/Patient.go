package hospital

type Patient struct {
	patientid     int    `json:"PatientID"  gorm:"primary_key"`
	doctorid      int    `json:"DoctorID"`
	patientname   string `json:"PatientName"`
	patientdob    string `json:"PatientDOB"`
	patientgender string `json:"PatientGender"`
}
