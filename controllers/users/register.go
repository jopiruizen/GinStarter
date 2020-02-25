package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-restapi/models"
	"go-restapi/services"
	"io/ioutil"
)

func serializeParam(ctx *gin.Context) (models.RegisterUserInput, error) {
	body := ctx.Request.Body
	value, err := ioutil.ReadAll(body)
	jsonParam := models.RegisterUserInput{}
	if err != nil {
		fmt.Println(err.Error())
		return models.RegisterUserInput{}, models.ErrMalformedParameter
	}
	json.Unmarshal(value, &jsonParam)
	return jsonParam, err
}

func Register(ctx *gin.Context) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("controllers.users.Register()")
	params, serializeErr := serializeParam(ctx)

	if serializeErr != nil {
		fmt.Println("Malformed Register POST JSON parameter")
		sendErrResponse(serializeErr, ctx)
		return
	}
	user, err := services.RegisterUser(params.User)

	if err != nil {
		sendErrResponse(err, ctx)
	} else {
		ctx.JSON(200, gin.H{
			"status_message": "Registration Successful",
			"user":           user.EncodeJSON(),
		})
	}
}
