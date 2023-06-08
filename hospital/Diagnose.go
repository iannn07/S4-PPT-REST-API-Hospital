package hospital

type Diagnose struct {
	diagnoseid           int    `json:"DiagnosisID" gorm:"primary_key"`
	patientid            int    `json:"PatientID"`
	doctorid             int    `json:"DoctorID"`
	diagnosisdate        string `json:"DiagnosisDate"`
	diagnosisdescription string `json:"DiagnosisDescription"`
}
