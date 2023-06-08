package hospital

type Doctor struct {
	Doctorid      int    `json:"DoctorID" gorm:"column:DoctorID;primaryKey"`
	Doctorname    string `json:"DoctorName"`
	Doctorlicense string `json:"DoctorLicense"`
}
