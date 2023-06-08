package diagnosehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/hospital"
)

func SelectAll(c *gin.Context) {
	var diagnoses []hospital.Diagnose
	hospital.DB.Find(&diagnoses)
	c.JSON(http.StatusOK, gin.H{"data": diagnoses})
}

func Create(c *gin.Context) {
	var input hospital.Diagnose
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	diagnose := hospital.Diagnose{diagnosisdate: input.diagnosisdate, diagnosisdescription: input.diagnosisdescription}
	hospital.DB.Create(&diagnose)

	c.JSON(http.StatusOK, gin.H{"data": diagnose})
}

func Read(c *gin.Context) {
	var diagnose hospital.Diagnose
	if err := hospital.DB.Where("id = ???", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diagnose})
}

func Update(c *gin.Context) {
	var diagnose hospital.Diagnose
	if err := diagnose.DB.Where("id = ???", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	var input hospital.Diagnose
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Hospital(&diagnose).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": diagnose})
}

func Delete(c *gin.Context) {
	var diagnose hospital.Diagnose
	if err := diagnose.DB.Where("id = ?", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Delete(&diagnose)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
