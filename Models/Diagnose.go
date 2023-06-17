package Models

type Diagnose struct {
	ID 			            int     `gorm:"column:diagnose_id;primaryKey;autoIncrement"			json:"diagnose_id"`
	DoctorID            	int     `gorm:"column:doctor_id"									json:"doctor_id"`
	PatientID            	int     `gorm:"column:patient_id"									json:"patient_id"`
	DiagnoseDate        	string  `gorm:"column:Diagnose_date;varchar(255)"					json:"Diagnose_date"`
	DiagnoseDescription 	string  `gorm:"column:Diagnose_description;varchar(255)"			json:"Diagnose_description"`
	Doctor               	Doctor  `gorm:"foreignKey:DoctorID" 								json:"doctor"`
	Patient              	Patient `gorm:"foreignKey:PatientID"								json:"patient"`
}

type DiagnoseInput struct {
	PatientID           int    	`json:"patient_id"`
	DoctorID            int    	`json:"doctor_id"`
	DiagnoseDate        string 	`json:"Diagnose_date"`
	DiagnoseDescription string 	`json:"Diagnose_description"`
}

type DiagnoseResponse struct {
	ID 			        int    `json:"diagnose_id"`
	PatientID           int    `json:"patient_id"`
	DoctorID            int    `json:"doctor_id"`
	DiagnoseDate        string `json:"Diagnose_date"`
	DiagnoseDescription string `json:"Diagnose_description"`
}

type DiagnoseListResponse struct {
	Diagnoses []DiagnoseResponse `json:"diagnoses"`
}
