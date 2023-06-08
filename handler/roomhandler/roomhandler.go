package roomhandler

import (
	"net/http"

	"HospitalFinpro/hospital"

	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var rooms []hospital.Room
	hospital.DB.Find(&rooms)
	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

func Create(c *gin.Context) {
	var input hospital.Room
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room := hospital.Room{Roomtype: input.Roomtype}
	hospital.DB.Create(&room)

	c.JSON(http.StatusOK, gin.H{"data": room})
}

func Read(c *gin.Context) {
	var room hospital.Room
	if err := hospital.DB.Where("RoomID = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room})
}

func Update(c *gin.Context) {
	var room hospital.Room
	if err := hospital.DB.Where("RoomID = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	var input hospital.Room
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Model(&room).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": room})
}

func Delete(c *gin.Context) {
	var room hospital.Room
	if err := hospital.DB.Where("RoomID = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}
	hospital.DB.Delete(&room)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
