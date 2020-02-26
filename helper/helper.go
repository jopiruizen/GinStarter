package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"go-restapi/models"
	"io/ioutil"
)

func BodyToJSONMap(ctx *gin.Context) (interface{}, error) {
	body := ctx.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var jsonMap interface{}
	json.Unmarshal(value, &jsonMap)
	log.Info("JSONBody: ", jsonMap)
	return jsonMap, err
}

/*
 * Helper Repository reusable function on Finding users in a list of users by email
 */
func SearchByEmail(email string, users []models.User) (models.User, error) {
	var target *models.User = nil

	log.Info(" ")
	log.Info("############# SearchByEmail() ", email)
	log.Info("#############  Users", users)

	for _, userData := range users {
		if email == userData.Email {
			target = &userData
			break
		}
	}

	log.Info("target:", target)
	if target == nil {
		return models.User{}, models.ErrNoRecordFound
	}
	return *target, nil
}
