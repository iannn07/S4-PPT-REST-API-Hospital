package hospital

type Payment struct {
	payid     int    `json:"PayID"  gorm:"primary_key"`
	patientid int    `json:"PatientID"`
	paytotal  string `json:"PayTotal"`
}
