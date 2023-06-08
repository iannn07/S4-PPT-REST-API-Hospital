package paymenthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iann07/S4-PPT-REST-API-Hospital.git/hospital"
)

func SelectAll(c *gin.Context) {
	var payments []hospital.Payment
	hospital.DB.Find(&payments)
	c.JSON(http.StatusOK, gin.H{"data": payments})
}

func Create(c *gin.Context) {
	var input hospital.Payment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment := hospital.Payment{paytotal: input.paytotal}
	hospital.DB.Create(&payment)

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

func Read(c *gin.Context) {
	var payment hospital.Payment
	if err := hospital.DB.Where("id = ???", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

func Update(c *gin.Context) {
	var payment hospital.Payment
	if err := hospital.DB.Where("id = ???", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	var input hospital.Payment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Hospital(&payment).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

func Delete(c *gin.Context) {
	var payment hospital.Payment
	if err := hospital.DB.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Delete(&payment)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
