package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-restapi/helper"
	"go-restapi/models"
)

func Register(context *gin.Context) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("controllers.users.Register()")
	jsonMap, err := helper.BodyToJSONMap(context)

	if err != nil {
		fmt.Println("Malformed POST JSON parameter")
	}

	user := models.User{}
	user.DecodeJSON(jsonMap)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("User Registered: User.Name: ", user.Name)
	fmt.Println("User Registered: User.Email: ", user.Email)
	fmt.Println("User Registered: User.Age: ", user.Age)

	status_message := "User Registered"

	errs := user.Validate()

	if errs != nil {
		status_message = "Validation Error!!"
	}

	context.JSON(200, gin.H{
		"status_message": status_message,
		"user": user.EncodeJSON(),
		"errors": errs,
	})
}