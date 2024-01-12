package routes

import (
	"service-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	r := router.Group("/user")
	{
		r.POST("/signup", controllers.UserSignup)
		// adding user location
		r.PUT("/addlocation/:user_id", controllers.GetLocation)
		// get all details of user by their id
		r.GET("/getdetails/:id", controllers.GetUserById)

		// login
		r.POST("/login", controllers.UserLogin)
		// catagory wise filtering
		r.GET("/plumbing", controllers.PlumpingCatagory)
		r.GET("/electritians", controllers.EletricianCatagory)
		r.GET("/painting", controllers.PaintingCatagory)
		r.GET("/cleaning", controllers.CleaningCatagory)
		r.GET("/cooking", controllers.CookingCatagory)
		r.GET("/others", controllers.OthersCatagory)

		//booking
		r.POST("/bookings/:user_id", controllers.BookingService)
		r.GET("/getbooking/:user_id", controllers.GetBooking)
		r.GET("/serivicers", controllers.GetAllServicer)

		r.GET("/popularservicer/:id", controllers.PopularServicers)
		r.POST("/addtosaved", controllers.AddToWistList)
		r.GET("/getsaved/:id", controllers.ShowSaved)

		r.DELETE("/saved/:saved_id", controllers.DeleteSaved)

		r.PUT("/update/:user_id", controllers.UpdateUser)
		r.PATCH("/success-razorpay/:booking_id",controllers.SuccessRazorPay)
	}

}
