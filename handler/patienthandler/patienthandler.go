package patienthandler

import (
	"net/http"

	"HospitalFinpro/hospital"

	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var patients []hospital.Patient
	hospital.DB.Find(&patients)
	c.JSON(http.StatusOK, gin.H{"data": patients})
}

func Create(c *gin.Context) {
	var input hospital.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient := hospital.Patient{PatientName: input.PatientName, PatientDOB: input.PatientDOB, PatientGender: input.PatientGender}
	hospital.DB.Create(&patient)

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

func Read(c *gin.Context) {
	var patient hospital.Patient
	if err := hospital.DB.Where("PatientID = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

func Update(c *gin.Context) {
	var patient hospital.Patient
	if err := hospital.DB.Where("PatientID = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	var input hospital.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Model(&patient).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

func Delete(c *gin.Context) {
	var patient hospital.Patient
	if err := hospital.DB.Where("PatientID = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Delete(&patient)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
