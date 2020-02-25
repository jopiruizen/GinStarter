package repository

import (
	"go-restapi/models"
)

type UserStore struct {
	usersList []models.User
}

func GetUserStore() *UserStore {
	store := UserStore{}
	store.usersList = []models.User{

		models.User{
			Name:  "Gon Freecs",
			Email: "gon.freecs@hxh.com",
			Age:   12,
		},

		models.User{
			Name:  "Killua Zoldyck",
			Email: "k.zoldyck@hxh.com",
			Age:   12,
		},

		models.User{
			Name:  "Larry David",
			Email: "ld@curb.com",
			Age:   75,
		},

		models.User{
			Name:  "Jerry Seinfeld",
			Email: "jerry@seinfeld.com",
			Age:   61,
		},
	}
	return &store
}

func (store *UserStore) FindByEmail(email string) (models.User, error) {

	var target *models.User = nil
	for _, userData := range store.usersList {
		if email == userData.Email {
			target = &userData
		}
	}

	if target == nil {
		return *target, models.ErrNoRecordFound
	}
	return *target, nil
}
