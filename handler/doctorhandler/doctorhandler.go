package doctorhandler

import (
	"net/http"

	"HospitalFinpro/hospital"

	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var doctors []hospital.Doctor
	hospital.DB.Find(&doctors)
	c.JSON(http.StatusOK, gin.H{"data": doctors})
}

func Create(c *gin.Context) {
	var input hospital.Doctor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	doctor := hospital.Doctor{Doctorname: input.Doctorname, Doctorlicense: input.Doctorlicense}
	hospital.DB.Create(&doctor)

	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

func Read(c *gin.Context) {
	var doctor hospital.Doctor
	if err := hospital.DB.Where("DoctorID = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

func Update(c *gin.Context) {
	var doctor hospital.Doctor
	if err := hospital.DB.Where("DoctorID = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	var input hospital.Doctor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Model(&doctor).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

func Delete(c *gin.Context) {
	var doctor hospital.Doctor
	if err := hospital.DB.Where("DoctorID = ?", c.Param("id")).First(&doctor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Delete(&doctor)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
