package hospital

type Payment struct {
	Payid     int    `json:"PayID"  gorm:"column:PayID;primaryKey"`
	Patientid int    `json:"PatientID"`
	Paytotal  string `json:"PayTotal"`
}
