package services

import (
	log "github.com/golang/glog"
	"go-restapi/models"
	"go-restapi/repository"
)

type Service struct {
	repository repository.IRepoSource
}

func NewService(repository repository.IRepoSource) Service {
	return Service{repository: repository}
}

func (service *Service) RegisterUser(user models.User) (models.User, error) {
	badInputError := user.Validate()
	/* DO REGISTRATION HERE */
	log.Info("")
	log.Info("")
	log.Info("User Registered: User.Name: ", user.Name)
	log.Info("User Registered: User.Email: ", user.Email)
	log.Info("User Registered: User.Age: ", user.Age)
	return user, badInputError
}

func (service *Service) Find(email string) (models.User, error) {
	log.Info("")
	log.Info("")
	log.Info("Finding User By Email: ", email)
	return service.repository.Find(email)
}
