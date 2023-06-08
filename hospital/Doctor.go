package hospital

type Doctor struct {
	doctorid      int    `json:"DoctorID" gorm:"primary_key"`
	doctorname    string `json:"DoctorName"`
	doctorlicense string `json:"DoctorLicense"`
}
