package DoctorController

import (
	"net/http"
	"HospitalFinpro/Models"
	"HospitalFinpro/Database"
	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var doctors []Models.Doctor
	Database.DB.Find(&doctors)

	var doctorResponses []Models.DoctorResponse
	for _, doctor := range doctors {
		doctorResponse := Models.DoctorResponse{
			ID:      doctor.ID,
			DoctorName:    doctor.DoctorName,
			DoctorLicense: doctor.DoctorLicense,
		}
		doctorResponses = append(doctorResponses, doctorResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": doctorResponses})
}

func Create(c *gin.Context) {
	var doctorInput Models.DoctorInput
	if err := c.ShouldBindJSON(&doctorInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor := Models.Doctor{
		DoctorName:    doctorInput.DoctorName,
		DoctorLicense: doctorInput.DoctorLicense,
	}

	if err := Database.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}

	doctorResponse := Models.DoctorResponse{
		ID:      doctor.ID,
		DoctorName:    doctor.DoctorName,
		DoctorLicense: doctor.DoctorLicense,
	}

	c.JSON(http.StatusOK, gin.H{"data": doctorResponse})
}

func Read(c *gin.Context) {
	var doctor Models.Doctor
	if err := Database.DB.Where("doctor_id = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DATA TIDAK DITEMUKAN!"})
		return
	}

	doctorResponse := Models.DoctorResponse{
		ID:      doctor.ID,
		DoctorName:    doctor.DoctorName,
		DoctorLicense: doctor.DoctorLicense,
	}

	c.JSON(http.StatusOK, gin.H{"data": doctorResponse})
}

func Update(c *gin.Context) {
	var doctor Models.Doctor
	if err := Database.DB.Where("doctor_id = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DATA TIDAK DITEMUKAN!"})
		return
	}

	var doctorInput Models.DoctorInput
	if err := c.ShouldBindJSON(&doctorInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor.DoctorName = doctorInput.DoctorName
	doctor.DoctorLicense = doctorInput.DoctorLicense

	if err := Database.DB.Save(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor"})
		return
	}

	doctorResponse := Models.DoctorResponse{
		ID:      doctor.ID,
		DoctorName:    doctor.DoctorName,
		DoctorLicense: doctor.DoctorLicense,
	}

	c.JSON(http.StatusOK, gin.H{"data": doctorResponse})
}

func Delete(c *gin.Context) {
	var doctor Models.Doctor
	if err := Database.DB.Where("doctor_id = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DATA TIDAK DITEMUKAN!"})
		return
	}

	if err := Database.DB.Delete(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
