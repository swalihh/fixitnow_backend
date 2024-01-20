package controllers

import (
	"fmt"
	"service-api/database"
	"service-api/helpers"
	"service-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	var input struct {
		User_Name string `json:"username"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Password  string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get data",
		})
		return
	}

	var user1 models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user1).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "email already exist",
		})
		return
	}

	pswd, err := helpers.HashPassword(input.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Password hashing error",
		})
		return
	}
	otp := helpers.GenerateOtp()

	helpers.SendOtp(strconv.Itoa(otp), input.Email)

	if err := database.DB.Create(&models.User{
		User_Name: input.User_Name,
		Email:     input.Email,
		Phone:     input.Phone,
		Password:  string(pswd),
		Otp:       otp,
	}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	var user models.User
	database.DB.Where("email = ?", input.Email).First(&user)

	fmt.Println(user.ID, otp, user.Otp)

	c.JSON(200, gin.H{
		"id":  user.ID,
		"otp": user.Otp,
	})

}

func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to find user",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetLocation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))
	var input struct {
		Location string `json:"location"`
	}
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get data",
		})
		return
	}
	if err := database.DB.Model(&models.User{}).Where("id = ?", id).Update("location", input.Location).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Updation error",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "successfully added location",
	})
}

func UserLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&input); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to  get data",
		})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Incorrect username and password",
		})
		return
	}
	if err := helpers.CheckPassword(user.Password, input.Password); err != nil {
		c.JSON(400, gin.H{
			"error": "Incorrect username and password",
		})
		return
	}
	c.JSON(200, gin.H{
		"id": user.ID,
	})
}

func PlumpingCatagory(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("service_catagory = ? AND status = ?", "Plumber", "Accepted").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get plumbers",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": servicers,
	})
}

func EletricianCatagory(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("service_catagory = ? AND status = ?", "Electrician", "Accepted").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get Electrician",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": servicers,
	})
}

func PaintingCatagory(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("service_catagory = ? AND status = ?", "Painting", "Accepted").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get Painting",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": servicers,
	})
}

func CleaningCatagory(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("service_catagory = ? AND status = ?", "Cleaning", "Accepted").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get Cleaning",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": servicers,
	})
}

func CookingCatagory(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("service_catagory = ? AND status = ?", "Cooking", "Accepted").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get Cooking",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": servicers,
	})
}

type Body struct {
	Catagory string `json:"catagory"`
}

func OthersCatagory(c *gin.Context) {
	var body Body
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get body",
		})
		return
	}
	fmt.Println("body:", body.Catagory)
	catagories := []string{"Plumber", "Electrician", "Cleaning", "Painting", "Cooking", "Others"}
	var servicers []models.Servicer
	if err := database.DB.Where("service_catagory NOT IN (?)", catagories).Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get catagory",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": servicers,
	})
}

func GetAllServicer(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("status=?", "Accepted").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "no accepted users",
		})
		return
	}
	c.JSON(200, gin.H{
		"accepted": servicers,
	})
}

type Datas struct {
	ID              int    `json:"servicer_id"`
	UserName        string `json:"username" gorm:"not null"`
	PhoneNumber     string `json:"phone" gorm:"not null"`
	FullName        string `json:"fullname"`
	Description     string `json:"description"`
	ServiceCatagory string `json:"servicecatagory"`
	Amount          int    `json:"amount"`
	Location        string `json:"location"`
	ServicerImage   string `json:"servicerimage"`
}

func PopularServicers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	var details []Datas
	if err := database.DB.Table("bookings").
		Select("servicers.id,servicers.user_name,servicers.phone_number,servicers.full_name,servicers.description,servicers.service_catagory,servicers.amount,servicers.amount,servicers.location,servicers.servicer_image").
		Joins("INNER JOIN servicers ON servicers.id=bookings.servicer_id").Where("bookings.user_id=? AND bookings.status=?", id, "Completed").Scan(&details).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find get data",
		})
		return
	}

	dm := make(map[Datas]bool)
	for _, datas := range details {
		dm[datas] = true
	}

	arr := []Datas{}

	for i := range dm {
		arr = append(arr, i)
	}

	c.JSON(200, gin.H{
		"data": arr,
	})

}

func AddToWistList(c *gin.Context) {
	usrId, _ := strconv.Atoi(c.Query("user_id"))
	servicerId, _ := strconv.Atoi(c.Query("servicer_id"))

	if err := database.DB.Create(&models.Saved{
		User_Id:     usrId,
		Servicer_Id: servicerId,
	}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to create data",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully added to saved",
	})
}

type Datass struct {
	Id              int    `json:"wistlist_id"`
	ID              int    `json:"servicer_id"`
	UserName        string `json:"username" gorm:"not null"`
	PhoneNumber     string `json:"phone" gorm:"not null"`
	FullName        string `json:"fullname"`
	Description     string `json:"description"`
	ServiceCatagory string `json:"servicecatagory"`
	Amount          int    `json:"amount"`
	Location        string `json:"location"`
	ServicerImage   string `json:"servicerimage"`
}

func ShowSaved(c *gin.Context) {
	usrId, _ := strconv.Atoi(c.Param("id"))

	var details []Datass

	if err := database.DB.Table("saveds").Select("saveds.id,servicers.id,servicers.user_name,servicers.phone_number,servicers.full_name,servicers.description,servicers.service_catagory,servicers.amount,servicers.amount,servicers.location,servicers.servicer_image").
		Joins("INNER JOIN servicers ON servicers.id=saveds.servicer_id").Where("saveds.user_id=?", usrId).Scan(&details).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find get data",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": details,
	})
}

func DeleteSaved(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("saved_id"))

	if err := database.DB.Where("id=?", id).Delete(&models.Saved{}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": "successfully deleted from saved",
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))

	type User struct {
		User_Name string `json:"username"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Location  string `json:"location"`
	}

	var input User
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get body",
		})
		return
	}
	data := map[string]interface{}{
		"user_name": input.User_Name,
		"email":     input.Email,
		"phone":     input.Phone,
		"location":  input.Location,
	}

	if err := database.DB.Model(&models.User{}).Where("id=?", id).Updates(data).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to  update datas",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully updated datas",
	})
}

func SuccessRazorPay(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("booking_id"))
	if err := database.DB.Model(&models.Booking{}).Where("id=?", id).Update("payment_status", "Completed").Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to  update data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully updated",
	})
}


func GetBookingDetails(c *gin.Context) {
	bookingId := c.Query("booking_id")
	userId := c.Query("user_id")

	var booking models.Booking
	if err :=database.DB.Where("id=? AND user_id=?",bookingId,userId).First(&booking).Error; err != nil {
		c.JSON(500,gin.H{
			"error" :"Failed to get data",
		})
		return
	}

	c.JSON(200,gin.H{
		"data":booking,
	})
}