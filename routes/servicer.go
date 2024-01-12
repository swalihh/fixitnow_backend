package routes

import (
	"service-api/controllers"

	"github.com/gin-gonic/gin"
)

func ServiceRouter(server *gin.Engine) {
	//grouping all routes with services with "r"
	r := server.Group("/servicer")

	r.POST("/signup", controllers.ServicerSignup)
	r.POST("/adddocuments/:id", controllers.AddDocuments)
	r.GET("/getdetails/:servicer_id", controllers.GetServicerDetails)
	r.POST("/login", controllers.ServicerLogin)
	r.GET("/getallbooking/:servicer_id", controllers.GetAllBookings)
	r.POST("/changestatus/:servicer_id", controllers.ChangeBookingStatus)
	r.PATCH("/addstarting-time/:booking_id", controllers.AddStartingTime)
	r.PATCH("/addending-time/:booking_id", controllers.AddEndingTime)
	r.GET("/get-rate/:booking_id", controllers.GetRateOfService)
	r.GET("/get-oneday-income/:servicer_id", controllers.GetOneDayIncome)
	r.GET("/get-onemonth-income/:servicer_id", controllers.GetOneMonthIncome)
	r.GET("/completed-count/:servicer_id", controllers.CountOfCompleted)
	r.GET("/accepted-count/:servicer_id", controllers.CountOfAccepted)
	r.PUT("/update/:servicer_id", controllers.UpdateServicer)
}
