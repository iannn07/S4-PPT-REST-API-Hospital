package patienthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/hospital"
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
	patient := hospital.Patient{patientname: input.patientname, patientdob: input.patientdob, patientgender: input.patientgender}
	hospital.DB.Create(&patient)

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

func Read(c *gin.Context) {
	var patient hospital.Patient
	if err := hospital.DB.Where("id = ???", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

func Update(c *gin.Context) {
	var patient hospital.Patient
	if err := hospital.DB.Where("id = ???", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	var input hospital.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Hospital(&patient).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

func Delete(c *gin.Context) {
	var patient hospital.Doctor
	if err := hospital.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Delete(&patient)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
