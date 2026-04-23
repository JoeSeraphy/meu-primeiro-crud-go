package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userID")
	r.GET("/getUserByEmail/:userEmail")
	r.POST("/createUser")
	r.PUT("/updateUser/:userID")
	r.DELETE("/deleteUser/:userID")
}
