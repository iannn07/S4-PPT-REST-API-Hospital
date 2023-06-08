package hospital

type Diagnose struct {
	Diagnoseid           int    `json:"DiagnosisID" gorm:"column:DiagnosisID;primaryKey"`
	Patientid            int    `json:"PatientID"`
	Doctorid             int    `json:"DoctorID"`
	Diagnosisdate        string `json:"DiagnosisDate"`
	Diagnosisdescription string `json:"DiagnosisDescription"`
}
