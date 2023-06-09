package hospital

type Doctor struct {
	DoctorID      int    `json:"DoctorID" gorm:"column:DoctorID;primaryKey"`
	DoctorName    string `json:"doctor_name"`
	DoctorLicense string `json:"doctor_license"`
}
