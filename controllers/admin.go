package controllers

import (
	"fmt"
	"os"
	"service-api/database"
	"service-api/helpers"
	"service-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input login
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get body",
		})
		return
	}

	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	fmt.Println(username, password, input)

	if username != input.Username || password != input.Password {
		c.JSON(400, gin.H{
			"error": "incorrect username or password",
		})
		return
	}

	token, err := helpers.GenerateJWT()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}

func GetPendingServicer(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("status=?", "Pending").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "no pending users",
		})
		return
	}
	c.JSON(200, gin.H{
		"pending": servicers,
	})
}

func GetAcceptedServicer(c *gin.Context) {
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

func GetRejectedServicer(c *gin.Context) {
	var servicers []models.Servicer
	if err := database.DB.Where("status=?", "Rejected").Find(&servicers).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "no rejected users",
		})
		return
	}
	c.JSON(200, gin.H{
		"rejected": servicers,
	})
}

func AcceptServicer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	if err := database.DB.Model(&models.Servicer{}).Where("id=?", id).Update("status", "Accepted").Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update status",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Updated status",
	})
}

func RejectServicer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	if err := database.DB.Model(&models.Servicer{}).Where("id=?", id).Update("status", "Rejected").Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to update status",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Updated status",
	})
}
