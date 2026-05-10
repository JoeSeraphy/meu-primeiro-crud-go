package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup,
	userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userID", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userID", userController.UpdateUser)
	r.DELETE("/deleteUser/:userID", userController.DeleteUser)
}
