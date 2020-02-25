package models

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

type RegisterUserInput struct {
	User User
}

type User struct {
	Name  string `validate:"nonzero"`
	Email string `validate:"nonzero"`
	Age   int    `validate:"min=13"`
}

func (user *User) Validate() error {
	if errs := validator.Validate(user); errs != nil {
		fmt.Println("Errors: ", errs)
		return ErrBadInput
	}
	return nil
}

func (user *User) EncodeJSON() string {
	data, error := json.Marshal(user)

	if error != nil {
		fmt.Println("Encode Error: ", error)
	}

	return string(data[:])
}
