package models

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

type User struct {
	Name  string `validate:"nonzero"`
	Email string `validate:"regexp=^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"`
	Age   int    `validate:"min=13"`
}

func (user *User) Validate() error {
	if errs := validator.Validate(user); errs != nil {
		fmt.Println("Errors: ", errs)
		return errs
	}
	return nil
}

func (user *User) EncodeJSON() []byte {
	data, error := json.Marshal(user)

	if error != nil {
		fmt.Println("Encode Error: ", error)
	}

	/*
		var data = make(map[string]interface{})
		data["name"] = user.Name
		data["email"] = user.Email
		data["age"] = user.Age

		//map[string]interface{}
	*/

	return data
}
