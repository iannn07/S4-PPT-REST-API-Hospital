package RoomController

import (
	"net/http"
	"HospitalFinpro/Models"
	"HospitalFinpro/Database"
	"github.com/gin-gonic/gin"
)

func SelectAll(c *gin.Context) {
	var rooms []Models.Room
	Database.DB.Find(&rooms)

	var roomResponses []Models.RoomResponse
	for _, room := range rooms {
		roomResponse := Models.RoomResponse{
			ID:    room.ID,
			PatientID: room.PatientID,
			RoomType:  room.RoomType,
		}
		roomResponses = append(roomResponses, roomResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": roomResponses})
}

func Create(c *gin.Context) {
	var roomInput Models.RoomInput
	if err := c.ShouldBindJSON(&roomInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := Models.Room{
		PatientID: roomInput.PatientID,
		RoomType:  roomInput.RoomType,
	}
	Database.DB.Create(&room)

	roomResponse := Models.RoomResponse{
		ID:    room.ID,
		PatientID: room.PatientID,
		RoomType:  room.RoomType,
	}

	c.JSON(http.StatusOK, gin.H{"data": roomResponse})
}

func Read(c *gin.Context) {
	var room Models.Room
	if err := Database.DB.Where("room_id = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	roomResponse := Models.RoomResponse{
		ID:    room.ID,
		PatientID: room.PatientID,
		RoomType:  room.RoomType,
	}

	c.JSON(http.StatusOK, gin.H{"data": roomResponse})
}

func Update(c *gin.Context) {
	var room Models.Room
	if err := Database.DB.Where("room_id = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var roomInput Models.RoomInput
	if err := c.ShouldBindJSON(&roomInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room.PatientID = roomInput.PatientID
	room.RoomType = roomInput.RoomType
	Database.DB.Save(&room)

	roomResponse := Models.RoomResponse{
		ID:    room.ID,
		PatientID: room.PatientID,
		RoomType:  room.RoomType,
	}

	c.JSON(http.StatusOK, gin.H{"data": roomResponse})
}

func Delete(c *gin.Context) {
	var room Models.Room
	if err := Database.DB.Where("room_id = ?", c.Param("id")).First(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	Database.DB.Delete(&room)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
