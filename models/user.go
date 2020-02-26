package models

import (
	"encoding/json"
	log "github.com/golang/glog"
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
		log.Error("Errors: ", errs.Error())
		return ErrBadInput
	}
	return nil
}

func (user *User) EncodeJSON() string {
	data, error := json.Marshal(user)

	if error != nil {
		log.Error("Encode Error: ", error.Error())
	}

	return string(data[:])
}
