package hospital

type Diagnose struct {
	DiagnoseID           int    `json:"DiagnosisID" gorm:"column:DiagnosisID;primaryKey"`
	PatientID            int    `json:"PatientID"`
	DoctorID             int    `json:"DoctorID"`
	DiagnosisDate        string `json:"DiagnosisDate"`
	DiagnosisDescription string `json:"DiagnosisDescription"`
}
