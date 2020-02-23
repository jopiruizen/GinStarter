package models

import (
	"fmt"
	"gopkg.in/validator.v2"
)

type User struct {
	Name string `validate:"nonzero"`
	Email string `validate:"regexp=^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"`
	Age int `validate:"min=13"`
}

func (user *User) DecodeJSON(json interface{}) {
	data := json.(map[string]interface{})
	userData := data["user"].(map[string]interface{})
	fmt.Println("models.User userdata: ", userData)
	user.Name = userData["name"].(string)
	user.Email = userData["email"].(string)
	user.Age = int(userData["age"].(float64))
	fmt.Println("models.User{}: ", user)
}

func (user *User) Validate() (error) {
	if errs := validator.Validate(user); errs != nil {
		fmt.Println("Errors: ", errs)
		return errs
	}
	return nil
}

func (user *User) EncodeJSON() (map[string]interface{}) {
	var data =  make(map[string]interface{})
	data["name"] = user.Name
	data["email"] = user.Email
	data["age"] = user.Age
	return data
}