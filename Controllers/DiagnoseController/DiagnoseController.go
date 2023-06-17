package DiagnoseController

import (
	"net/http"
	"HospitalFinpro/Models"
	"HospitalFinpro/Database"
	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var diagnoses []Models.Diagnose
	Database.DB.Find(&diagnoses)

	var diagnoseResponses []Models.DiagnoseResponse
	for _, diagnose := range diagnoses {
		diagnoseResponse := Models.DiagnoseResponse{
			ID:          			diagnose.ID,
			DiagnoseDate:        	diagnose.DiagnoseDate,
			DiagnoseDescription: 	diagnose.DiagnoseDescription,
			DoctorID:            	diagnose.DoctorID,
			PatientID:            	diagnose.PatientID,
			
		}
		diagnoseResponses = append(diagnoseResponses, diagnoseResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": diagnoseResponses})
}

func Create(c *gin.Context) {
	var diagnoseInput Models.DiagnoseInput
	if err := c.ShouldBindJSON(&diagnoseInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diagnose := Models.Diagnose{
		DiagnoseDate:       	diagnoseInput.DiagnoseDate,
		DiagnoseDescription: 	diagnoseInput.DiagnoseDescription,
		DoctorID:             	diagnoseInput.DoctorID,
		PatientID:            	diagnoseInput.PatientID,
	}

	if err := Database.DB.Create(&diagnose).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create diagnose"})
		return
	}

	diagnoseResponse := Models.DiagnoseResponse{
		ID:         			 	diagnose.ID,
		DiagnoseDate:       		diagnoseInput.DiagnoseDate,
		DiagnoseDescription: 		diagnose.DiagnoseDescription,
		PatientID:            		diagnose.PatientID,
		DoctorID:             		diagnose.DoctorID,
	}

	c.JSON(http.StatusOK, gin.H{"data": diagnoseResponse})
}

func Read(c *gin.Context) {
	var diagnose Models.Diagnose
	if err := Database.DB.Where("diagnose_id = ?", c.Param("id")).Preload("Patient").Preload("Doctor").First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	diagnoseResponse := Models.DiagnoseResponse{
		ID:          			diagnose.ID,
		DiagnoseDate:        	diagnose.DiagnoseDate,
		DiagnoseDescription: 	diagnose.DiagnoseDescription,
		DoctorID:             	diagnose.DoctorID,
		PatientID:            	diagnose.PatientID,
	}

	c.JSON(http.StatusOK, gin.H{"data": diagnoseResponse})
}

func Update(c *gin.Context) {
	var diagnose Models.Diagnose
	if err := Database.DB.Where("diagnose_id = ?", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var diagnoseInput Models.DiagnoseInput
	if err := c.ShouldBindJSON(&diagnoseInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diagnose.DiagnoseDate = diagnoseInput.DiagnoseDate
	diagnose.DiagnoseDescription = diagnoseInput.DiagnoseDescription
	diagnose.DoctorID = diagnoseInput.DoctorID
	diagnose.PatientID = diagnoseInput.PatientID

	if err := Database.DB.Save(&diagnose).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update diagnose"})
		return
	}

	diagnoseResponse := Models.DiagnoseResponse{
		ID:          		  diagnose.ID,
		DiagnoseDate:         diagnose.DiagnoseDate,
		DiagnoseDescription:  diagnose.DiagnoseDescription,
		DoctorID:             diagnose.DoctorID,
		PatientID:            diagnose.PatientID,
	}

	c.JSON(http.StatusOK, gin.H{"data": diagnoseResponse})
}

func Delete(c *gin.Context) {
	var diagnose Models.Diagnose
	if err := Database.DB.Where("diagnose_id = ?", c.Param("id")).First(&diagnose).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := Database.DB.Delete(&diagnose).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete diagnose"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
