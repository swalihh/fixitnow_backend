package routes

import (
	"service-api/controllers"
	"service-api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(server *gin.Engine) {
	r := server.Group("/admin")
	r.POST("/login", controllers.AdminLogin)
	r.GET("/pendingservicer", middleware.AdminAuthentication, controllers.GetPendingServicer)
	r.GET("/acceptedservicer", middleware.AdminAuthentication, controllers.GetAcceptedServicer)
	r.GET("/rejectedservicer", middleware.AdminAuthentication, controllers.GetRejectedServicer)
	r.PATCH("/acceptservicer/:servicer_id", middleware.AdminAuthentication, controllers.AcceptServicer)
	r.PATCH("/rejectservicer/:servicer_id", middleware.AdminAuthentication, controllers.RejectServicer)
}
