package Models

type Patient struct {
	ID 			    int    `gorm:"column:patient_id;primaryKey;autoIncrement"		json:"patient_id"`
	DoctorID     	int    `gorm:"column:doctor_id"									json:"doctor_id"`
	PatientName   	string `gorm:"column:patient_name;varchar(255)"					json:"patient_name`
	PatientDOB    	string `gorm:"column:patient_dob;varchar(255)"					json:"patient_dob"`
	PatientGender 	string `gorm:"column:patient_gender;varchar(255)"				json:"patient_gender`
	Doctor        	Doctor `gorm:"foreignKey:DoctorID"								json:"doctor"`
}

type PatientInput struct {
	DoctorID      int    `json:"doctor_id"`
	PatientName   string `json:"patient_name"`
	PatientDOB    string `json:"patient_dob"`
	PatientGender string `json:"patient_gender"`
}

type PatientResponse struct {
	ID     		  int    `json:"patient_id"`
	DoctorID      int    `json:"doctor_id"`
	PatientName   string `json:"patient_name"`
	PatientDOB    string `json:"patient_dob"`
	PatientGender string `json:"patient_gender"`
}

type PatientListResponse struct {
	Patients []PatientResponse `json:"patients"`
}
