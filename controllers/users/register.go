package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-restapi/models"
	"go-restapi/services"
	"io/ioutil"
)

type RegisterJSONParam struct {
	user models.User
}

func serializeParam(ctx *gin.Context) (RegisterJSONParam, error) {
	body := ctx.Request.Body
	value, err := ioutil.ReadAll(body)
	jsonParam := RegisterJSONParam{}
	if err != nil {
		fmt.Println(err.Error())
		return RegisterJSONParam{}, err
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
	}
	user, statusMessage, errs := services.RegisterUser(params.user)
	ctx.JSON(200, gin.H{
		"status_message": statusMessage,
		"user":           user.EncodeJSON(),
		"errors":         errs,
	})
}
