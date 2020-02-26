package static

import (
	log "github.com/golang/glog"
	"go-restapi/helper"
	"go-restapi/models"
)

type StaticSource struct {
	usersList []models.User
}

func (src *StaticSource) LoadSource() {

	log.Info("StaticSource.LoadSource()")
	src.usersList = []models.User{
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

func (src *StaticSource) Find(email string) (models.User, error) {
	return helper.SearchByEmail(email, src.usersList)
}
