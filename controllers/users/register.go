package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"go-restapi/models"
	"go-restapi/repository"
	"go-restapi/repository/commons"
	"go-restapi/services"
	"io/ioutil"
)

func serializeParam(ctx *gin.Context) (models.RegisterUserInput, error) {
	body := ctx.Request.Body
	value, err := ioutil.ReadAll(body)
	jsonParam := models.RegisterUserInput{}
	if err != nil {
		log.Error(err.Error())
		return models.RegisterUserInput{}, models.ErrMalformedParameter
	}
	json.Unmarshal(value, &jsonParam)
	log.Info("SerializeParam: ", jsonParam)
	return jsonParam, err
}

func Register(ctx *gin.Context) {
	log.Info("")
	log.Info("")
	log.Info("controllers.users.Register()")
	params, serializeErr := serializeParam(ctx)

	if serializeErr != nil {
		log.Error("Malformed Register POST JSON parameter", serializeErr.Error())
		sendErrResponse(serializeErr, ctx)
		return
	}

	var source = repository.NewRepoSource(commons.SOURCE_TYPE_STATIC)
	var service = services.NewService(source)
	user, err := service.RegisterUser(params.User)

	if err != nil {
		sendErrResponse(err, ctx)
	} else {
		ctx.JSON(200, gin.H{
			"status_message": "Registration Successful",
			"user":           user.EncodeJSON(),
		})
	}
}
