package services

import (
	log "github.com/golang/glog"
	"go-restapi/models"
	"go-restapi/repository"
	"go-restapi/repository/commons"
)

func RegisterUser(user models.User) (models.User, error) {
	badInputError := user.Validate()
	/* DO REGISTRATION HERE */

	log.Info("")
	log.Info("")
	log.Info("User Registered: User.Name: ", user.Name)
	log.Info("User Registered: User.Email: ", user.Email)
	log.Info("User Registered: User.Age: ", user.Age)
	return user, badInputError
}

func FindUserByEmail(email string) (models.User, error) {
	log.Info("")
	log.Info("")
	log.Info("Finding User By Email: ", email)
	source := repository.NewRepoSource(commons.SOURCE_TYPE_STATIC)
	return source.Find(email)
}

func FindUserByEmailOnFile(email string) (models.User, error) {
	log.Info("")
	log.Info("")
	log.Info("Finding User By Email: ", email)
	source := repository.NewRepoSource(commons.SOURCE_TYPE_FILE)
	return source.Find(email)
}
