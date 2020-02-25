package repository

import (
	"go-restapi/models"
)

type StaticSource struct{}

func (src *StaticSource) LoadSource() []models.User {
	return []models.User{
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
}
