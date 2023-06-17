package PaymentController

import (
	"net/http"
	"HospitalFinpro/Models"
	"HospitalFinpro/Database"
	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var payments []Models.Payment
	Database.DB.Find(&payments)

	var paymentResponses []Models.PaymentResponse
	for _, payment := range payments {
		paymentResponse := Models.PaymentResponse{
			ID:     payment.ID,
			PatientID: payment.PatientID,
			PayTotal:  payment.PayTotal,
		}
		paymentResponses = append(paymentResponses, paymentResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentResponses})
}

func Create(c *gin.Context) {
	var paymentInput Models.PaymentInput
	if err := c.ShouldBindJSON(&paymentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := Models.Payment{
		PatientID: paymentInput.PatientID,
		PayTotal:  paymentInput.PayTotal,
	}
	if err := Database.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	paymentResponse := Models.PaymentResponse{
		ID:     payment.ID,
		PatientID: payment.PatientID,
		PayTotal:  payment.PayTotal,
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentResponse})
}

func Read(c *gin.Context) {
	var payment Models.Payment
	if err := Database.DB.Where("pay_id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	paymentResponse := Models.PaymentResponse{
		ID:     payment.ID,
		PatientID: payment.PatientID,
		PayTotal:  payment.PayTotal,
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentResponse})
}

func Update(c *gin.Context) {
	var payment Models.Payment
	if err := Database.DB.Where("pay_id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var paymentInput Models.PaymentInput
	if err := c.ShouldBindJSON(&paymentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.PatientID = paymentInput.PatientID
	payment.PayTotal = paymentInput.PayTotal

	if err := Database.DB.Save(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		return
	}

	paymentResponse := Models.PaymentResponse{
		ID:     payment.ID,
		PatientID: payment.PatientID,
		PayTotal:  payment.PayTotal,
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentResponse})
}

func Delete(c *gin.Context) {
	var payment Models.Payment
	if err := Database.DB.Where("pay_id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := Database.DB.Delete(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
