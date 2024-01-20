package controllers

import (
	"database/sql"
	"fmt"
	"service-api/database"
	"service-api/helpers"
	"service-api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ServicerSignup(c *gin.Context) {
	var input struct {
		UserName    string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone"`
		Password    string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	var servicer1 models.Servicer
	if err := database.DB.Where("email = ?", input.Email).First(&servicer1).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "email already exist",
		})
		return
	}

	otp := helpers.GenerateOtp()

	otpstr := strconv.Itoa(otp)

	helpers.SendOtp(otpstr, input.Email)

	password, _ := helpers.HashPassword(input.Password)

	if err := database.DB.Create(&models.Servicer{
		UserName:    input.UserName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    string(password),
		OTP:         otpstr,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var servicer models.Servicer

	database.DB.Last(&servicer)

	c.JSON(200, gin.H{
		"id":  servicer.ID,
		"otp": servicer.OTP,
	})
}

func AddDocuments(c *gin.Context) {
	var input struct {
		FullName             string `json:"fullname"`
		Description          string `json:"description"`
		ServiceCatagory      string `json:"servicecatagory"`
		VerificationDocument string `json:"verificationdocument"`
		Amount               int    `json:"amount"`
		Location             string `json:"location"`
		ServicerImage        string `json:"servicerimage"`
		ServicerDocument     string `json:"servicerdocument"`
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Binding error",
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Model(&models.Servicer{}).Where("id=?", id).Updates(map[string]interface{}{"full_name": input.FullName,
		"description": input.Description, "service_catagory": input.ServiceCatagory,
		"verification_document": input.VerificationDocument, "amount": input.Amount, "location": input.Location,
		"servicer_image": input.ServicerImage, "servicer_document": input.ServicerDocument, "status": "Pending"}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "updation error",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "created new servicer",
	})

}

func GetServicerDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	var servicer models.Servicer
	if err := database.DB.First(&servicer, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find servicer",
		})
		return
	}
	c.JSON(200, gin.H{
		"servicer": servicer,
	})
}

func ServicerLogin(c *gin.Context) {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&login); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	var servicer models.Servicer
	if err := database.DB.Where("email = ?", login.Email).First(&servicer).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "incorrect eamail and password",
		})
		return
	}
	if err := helpers.CheckPassword(servicer.Password, login.Password); err != nil {
		c.JSON(400, gin.H{
			"error": "incorrect eamail and password",
		})
		return
	}
	c.JSON(200, gin.H{
		"id": servicer.ID,
	})

}

type BookingDetails struct {
	Id            int    `json:"id"`
	BuildingName  string `json:"buildingname"`
	City          string `json:"city"`
	Road          string `json:"road"`
	Phone         string `json:"phone"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Description   string `json:"description"`
	ServiceAmount int    `json:"serviceamount"`
	Status        string `json:"status"`
	User_Name     string `json:"username"`
	Revenue       int64  `json:"revenue"`
}

func GetAllBookings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))

	var booking []BookingDetails

	if err := database.DB.Table("bookings").Select("bookings.id,bookings.building_name,bookings.city,bookings.road,bookings.phone,bookings.date,bookings.time,bookings.description,bookings.service_amount,bookings.status,users.user_name,bookings.revenue").
		Joins("INNER JOIN users ON users.id=bookings.user_id").Where("bookings.servicer_id=?", id).Scan(&booking).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": booking,
	})
}

func ChangeBookingStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	var input struct {
		BookingID int    `json:"bookingid"`
		Status    string `json:"status"`
	}
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get body",
		})
		return
	}

	fmt.Println(input, id)

	var booking models.Booking
	if err := database.DB.Where("id = ? AND servicer_id = ?", input.BookingID, id).Find(&booking).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "this servicer not access this booking",
		})
		return
	}
	if err := database.DB.Model(&models.Booking{}).Where("id = ?", input.BookingID).Update("status", input.Status).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to update data",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": "successfully updated status",
	})
}

func AddStartingTime(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("booking_id"))
	time := time.Now()
	if err := database.DB.Model(models.Booking{}).Where("id = ?", id).Update("starting_time", time).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update time",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":       "successfully updated starting time",
		"starting-time": time,
	})

}

func AddEndingTime(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("booking_id"))
	time := time.Now()
	if err := database.DB.Model(models.Booking{}).Where("id = ?", id).Update("ending_time", time).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update time",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":     "successfully updated ending time",
		"ending-time": time,
	})

}

func GetRateOfService(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("booking_id"))

	var booking models.Booking
	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get booking details",
		})
		return
	}
	hour := float64(booking.EndingTime.Unix()-booking.StartingTime.Unix()) / 60 / 60
	rate := hour * float64(booking.ServiceAmount)

	database.DB.Model(models.Booking{}).Where("id=?", id).Update("revenue", int64(rate))

	c.JSON(200, gin.H{
		"rate":       int(rate),
		"start-time": booking.StartingTime,
		"end-time":   booking.EndingTime,
	})
}

func GetOneDayIncome(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	current := time.Now()
	today := current.Format("2006-01-02")
	var revenue int
	if err := database.DB.Table("bookings").Select("SUM(revenue)").Where("date = ? AND status = ? AND servicer_id = ?", today, "Completed", id).Scan(&revenue).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to find revenue",
		})
		return
	}
	c.JSON(200, gin.H{
		"revenue": revenue,
	})
}

func GetOneMonthIncome(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))

	current := time.Now().AddDate(0, 0, 1)
	today := current.Format("2006-01-02")

	before30 := time.Now().AddDate(0, 0, -30)
	before := before30.Format("2006-01-02")

	fmt.Println(before)

	var revenue sql.NullInt64
	if err := database.DB.Table("bookings").Select("SUM(revenue)").Where("date BETWEEN ? AND ? AND status = ? AND servicer_id = ?", before, today, "Completed", id).Scan(&revenue).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to find revenue",
		})
		return
	}
	c.JSON(200, gin.H{
		"revenue": revenue.Int64,
	})
}

func CountOfCompleted(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))

	var count int
	if err := database.DB.Table("bookings").Select("COUNT(*)").Where("servicer_id = ? AND status = ?", id, "Completed").Scan(&count).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get count",
		})
		return
	}

	c.JSON(200, gin.H{
		"count": count,
	})
}

func CountOfAccepted(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))

	var count int
	if err := database.DB.Table("bookings").Select("COUNT(*)").Where("servicer_id = ? AND status = ? OR status = ?", id, "Accepted", "Completed").Scan(&count).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get count",
		})
		return
	}

	c.JSON(200, gin.H{
		"count": count,
	})
}

func UpdateServicer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))

	type Servicer struct {
		UserName        string `json:"username"`
		Email           string `json:"email"`
		PhoneNumber     string `json:"phone"`
		FullName        string `json:"fullname"`
		Description     string `json:"description"`
		ServiceCatagory string `json:"servicecatagory"`
		Amount          int    `json:"amount"`
		Location        string `json:"location"`
		ServicerImage   string `json:"servicerimage"`
	}

	var input Servicer

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get body please check json",
		})
		return
	}

	data := map[string]interface{}{
		"user_name":        input.UserName,
		"email":            input.Email,
		"phone_number":     input.PhoneNumber,
		"full_name":        input.FullName,
		"description":      input.Description,
		"service_catagory": input.ServiceCatagory,
		"amount":           input.Amount,
		"location":         input.Location,
		"servicer_image":   input.ServicerImage,
	}

	if err := database.DB.Model(&models.Servicer{}).Where("id=?", id).Updates(data).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to update datas",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully updated data",
	})

}
