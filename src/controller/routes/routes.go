package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userID", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userID", controller.UpdateUser)
	r.DELETE("/deleteUser/:userID", controller.DeleteUser)
}
