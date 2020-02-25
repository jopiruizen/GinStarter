package services

import (
	"fmt"
	"go-restapi/models"
	"go-restapi/repository"
)

func RegisterUser(user models.User) (models.User, error) {
	badInputError := user.Validate()

	/* DO REGISTRATION HERE */
	fmt.Println("")
	fmt.Println("")
	fmt.Println("User Registered: User.Name: ", user.Name)
	fmt.Println("User Registered: User.Email: ", user.Email)
	fmt.Println("User Registered: User.Age: ", user.Age)

	return user, badInputError
}

func FindUserByEmail(email string) (models.User, error) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Finding User By Email: ", email)
	store := repository.GetUserStore()
	return store.FindByEmail(email)
}
