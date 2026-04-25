package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect filds, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
