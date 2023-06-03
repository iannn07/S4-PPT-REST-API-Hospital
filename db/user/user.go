package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	patientid     int    `json:"PatientID"`
	doctorid      int    `json:"DoctorID"`
	patientname   string `json:"PatientName"`
	patientdob    string `json:"PatientDOB"`
	patientgender string `json:"PatientGender"`
}
type Doctor struct {
	doctorid      int    `json:"DoctorID"`
	doctorname    string `json:"DoctorName"`
	doctorlicense string `json:"DoctorLicense"`
}
type Diagnose struct {
	diagnoseid           int    `json:"DiagnosisID"`
	patientid            int    `json:"PatientID"`
	doctorid             int    `json:"DoctorID"`
	diagnosisdate        string `json:"DiagnosisDate"`
	diagnosisdescription string `json:"DiagnosisDescription"`
}
type Room struct {
	roomid    int    `json:"RoomID"`
	patientid int    `json:"PatientID"`
	roomtype  string `json:"RoomType"`
}
type Payment struct {
	payid     int    `json:"PayID"`
	patientid int    `json:"PatientID"`
	paytotal  string `json:"PayTotal"`
}

func cekError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:83306)/finpro_ppt")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// PATIENT
func SelectAll() ([]Patient, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM patient")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var patients []Patient
	for result.Next() {
		var patient Patient
		err := result.Scan(&patient.patientid, &patient.doctorid, &patient.patientname, &patient.patientdob, &patient.patientgender)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return patients, nil
}

func SelectByID(id int) (Patient, error) {
	db, err := connect()
	if err != nil {
		return Patient{}, err
	}
	defer db.Close()

	result := db.QueryRow("SELECT * FROM patient Where PatientID = ?", id)

	var patient Patient
	err = result.Scan(&patient.patientid, &patient.doctorid, &patient.patientname, &patient.patientdob, &patient.patientgender)
	if err != nil {
		return Patient{}, err
	}
	return patient, nil
}

func Insert(patient Patient) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO patient (PatientID, DoctorID, PatientName, PatientDOB, PatientGender) VALUES (11111111111, 22222222222, Dora Kalifa Dharmawan, 01/01/2001, Female)",
		patient.patientid, patient.doctorid, patient.patientname, patient.patientdob, patient.patientgender)
	if err != nil {
		return err
	}
	return nil
}

func Update(patient Patient) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE patient SET DoctorID= 11111111111, PatientName= Dora Khalifa Dharmawan, PatientDOB= 01/01/2001, PatientGender= Female WHERE PatientID= 22222222222",
		patient.doctorid, patient.patientname, patient.patientdob, patient.patientgender, patient.patientid)
	if err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM patient WHERE PatientID= 22222222222", id)
	if err != nil {
		return err
	}
	return nil
}

// DOCTOR
func SelectAllDoctors() ([]Doctor, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM doctor")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var doctors []Doctor
	for result.Next() {
		var doctor Doctor
		err := result.Scan(&doctor.doctorid, &doctor.doctorname, &doctor.doctorlicense)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return doctors, nil
}

func InsertDoctor(doctor Doctor) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO doctor (DoctorID, DoctorName, DoctorLicense) VALUES (?, ?, ?)",
		doctor.doctorid, doctor.doctorname, doctor.doctorlicense)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDoctor(doctor Doctor) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE doctor SET DoctorName=?, DoctorLicense=? WHERE DoctorID=?",
		doctor.doctorname, doctor.doctorlicense, doctor.doctorid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDoctor(doctorID int) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM doctor WHERE DoctorID=?", doctorID)
	if err != nil {
		return err
	}
	return nil
}

// DIAGNOSES
func SelectAllDiagnoses() ([]Diagnose, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM diagnose")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var diagnoses []Diagnose
	for result.Next() {
		var diagnose Diagnose
		err := result.Scan(&diagnose.diagnoseid, &diagnose.patientid, &diagnose.doctorid, &diagnose.diagnosisdate, &diagnose.diagnosisdescription)
		if err != nil {
			return nil, err
		}
		diagnoses = append(diagnoses, diagnose)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return diagnoses, nil
}

func InsertDiagnoses(diagnose Diagnose) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO diagnose (DiagnoseID, PatientID, DoctorID, DiagnosisDate, DiagnosisDescription) VALUES (?, ?, ?, ?, ?)",
		diagnose.diagnoseid, diagnose.patientid, diagnose.doctorid, diagnose.diagnosisdate, diagnose.diagnosisdescription)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDiagnose(diagnose Diagnose) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE diagnose SET PatientID=?, DoctorID=?, DiagnosisDate=?, DiagnosisDescription=? WHERE DiagnoseID=?",
		diagnose.patientid, diagnose.doctorid, diagnose.diagnosisdate, diagnose.diagnosisdescription, diagnose.diagnoseid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDiagnose(diagnoseID int) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM diagnose WHERE DiagnoseID=?", diagnoseID)
	if err != nil {
		return err
	}

	return nil
}

// ROOMS
func SelectAllRooms() ([]Room, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM room")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var rooms []Room
	for result.Next() {
		var room Room
		err := result.Scan(&room.roomid, &room.patientid, &room.roomtype)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}

func InsertRoom(room Room) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO room (RoomID, PatientID, RoomType) VALUES (?, ?, ?)",
		room.roomid, room.patientid, room.roomtype)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRoom(room Room) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE room SET PatientID=?, RoomType=? WHERE RoomID=?",
		room.patientid, room.roomtype, room.roomid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRoom(roomID int) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM room WHERE RoomID=?", roomID)
	if err != nil {
		return err
	}
	return nil
}

// PAYMENT
func SelectAllPayments() ([]Payment, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM payment")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var payments []Payment
	for result.Next() {
		var payment Payment
		err := result.Scan(&payment.payid, &payment.patientid, &payment.paytotal)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return payments, nil
}

func InsertPayment(payment Payment) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO payment (PayID, PatientID, PayTotal) VALUES (?, ?, ?)",
		payment.payid, payment.patientid, payment.paytotal)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePayment(payment Payment) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE payment SET PatientID=?, PayTotal=? WHERE PayID=?",
		payment.patientid, payment.paytotal, payment.payid)
	if err != nil {
		return err
	}
	return nil
}

func DeletePayment(paymentID int) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM payment WHERE PayID=?", paymentID)
	if err != nil {
		return err
	}
	return nil
}


func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	var patient Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = InsertPatient(patient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
	}
}
func ReadPatientHandler(w http.ResponseWriter, r *http.Request) {
	patients, err := SelectAllPatients()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(patients)
}
func UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	var patient Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = UpdatePatient(patient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	patientID := //parse tha patient ID from the request parameter or body

	err := DeletePatient(patientID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
