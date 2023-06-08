package diagnosehandler

import (
	"HospitalFinpro/hospital"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var diagnose []hospital.Diagnose
	hospital.DB.Find(&diagnose)
	c.JSON(http.StatusOK, gin.H{"data": diagnose})
}

func Create(c *gin.Context) {
	var input hospital.Diagnose
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	diagnose := hospital.Diagnose{Diagnosisdate: input.Diagnosisdate, Diagnosisdescription: input.Diagnosisdescription}
	hospital.DB.Create(&diagnose)

	c.JSON(http.StatusOK, gin.H{"data": diagnose})
}

func Read(c *gin.Context) {
	var diagnose hospital.Diagnose
	if err := hospital.DB.Where("DiagnosisID = ?", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diagnose})
}

func Update(c *gin.Context) {
	var diagnose hospital.Diagnose
	if err := hospital.DB.Where("DiagnosisID = ?", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	var input hospital.Diagnose
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Model(&diagnose).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": diagnose})
}

func Delete(c *gin.Context) {
	var diagnose hospital.Diagnose
	if err := hospital.DB.Where("DiagnosisID = ?", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Delete(&diagnose)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
