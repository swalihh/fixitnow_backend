package controllers

import (
	"service-api/database"
	"service-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookingService(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))
	var input models.Booking
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get body",
		})
		return
	}
	if err := database.DB.Create(&models.Booking{
		BuildingName:  input.BuildingName,
		City:          input.City,
		Road:          input.Road,
		Phone:         input.Phone,
		Date:          input.Date,
		Time:          input.Time,
		Description:   input.Description,
		Servicer_id:   input.Servicer_id,
		ServiceAmount: input.ServiceAmount,
		User_Id:       id,
		Status:        "Pending",
	}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to add data",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Successfully booked servicer",
	})
}

type BookingDetail struct {
	Date            string `json:"date"`
	Time            string `json:"time"`
	Description     string `json:"description"`
	FullName        string `json:"fullname"`
	UserName        string `json:"username"`
	PhoneNumber     string `json:"phone"`
	ServiceCatagory string `json:"servicecatagory"`
	ServiceAmount   int    `json:"serviceamount"`
	Status          string `json:"status"`
	ServicerImage   string `json:"servicerimage"`
	Location             string `json:"location"`
}

func GetBooking(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))

	var booking []BookingDetail

	if err := database.DB.Table("bookings").
		Select("bookings.date,bookings.time,bookings.description,servicers.full_name,servicers.user_name,servicers.phone_number,servicers.service_catagory,bookings.service_amount,bookings.status,servicers.servicer_image,servicers.location").
		Joins("INNER JOIN servicers ON servicers.id=bookings.servicer_id").Where("bookings.user_id=?", id).Find(&booking).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": booking,
	})
}
