package repository

import (
	"fmt"
	"go-restapi/models"
)

type UserStore struct {
	usersList []models.User
}

type ISourceLoader interface {
	LoadSource() []models.User
}

func GetUserStore(sourceLoader ISourceLoader) UserStore {
	store := UserStore{}
	store.usersList = sourceLoader.LoadSource()
	return store
}

func (store *UserStore) FindByEmail(email string) (models.User, error) {
	var target *models.User = nil
	for _, userData := range store.usersList {
		if email == userData.Email {
			target = &userData
			break
		}
	}
	fmt.Println("FindByEmail() ", email)
	fmt.Println("target:", target)
	if target == nil {
		return models.User{}, models.ErrNoRecordFound
	}
	return *target, nil
}
