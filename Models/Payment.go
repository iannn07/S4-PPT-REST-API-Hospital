package Models

type Payment struct {
	ID     		int     `gorm:"column:pay_id;primaryKey;autoIncrement"			json:"payment_id"`
	PatientID 	int     `gorm:"column:patient_id"								json:"patient_id"`
	PayTotal  	string  `gorm:"column:pay_total;varchar(255)"					json:"payment_total"`
	Patient   	Patient `gorm:"foreignKey:PatientID"							json:"patient"`
}

type PaymentInput struct {
	PatientID int    `json:"patient_id"`
	PayTotal  string `json:"pay_total"`
}

type PaymentResponse struct {
	ID     		int    `json:"pay_id"`
	PatientID 	int    `json:"patient_id"`
	PayTotal  	string `json:"pay_total"`
}

type PaymentListResponse struct {
	Payments []PaymentResponse `json:"payments"`
}
