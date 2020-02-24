package services

import (
	"fmt"
	"go-restapi/models"
)

func RegisterUser( user models.User) (models.User, string, error) {
	validationErrors := user.Validate()
	statusMessage := "User Successfully Registered!"
	if validationErrors != nil {
		statusMessage = "Validation Errors!!!"
	}
	/* DO REGISTRATION HERE */
	fmt.Println("")
	fmt.Println("")
	fmt.Println("User Registered: User.Name: ", user.Name)
	fmt.Println("User Registered: User.Email: ", user.Email)
	fmt.Println("User Registered: User.Age: ", user.Age)
	return user, statusMessage, validationErrors
}