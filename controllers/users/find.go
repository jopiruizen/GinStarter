package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"go-restapi/services"
)

func Find(ctx *gin.Context) {
	log.Info("")
	log.Info("")
	log.Info("hello")
	log.Info("controllers.users.Find()")
	params, serializeErr := serializeParam(ctx)
	fmt.Println("User param: ", params)
	fmt.Println("Error:?", serializeErr)
	if serializeErr != nil {
		fmt.Println("Malformed Register POST JSON parameter")
		sendErrResponse(serializeErr, ctx)
		return
	}

	user, findError := services.FindUserByEmail(params.User.Email)
	if findError != nil {
		sendErrResponse(findError, ctx)
	} else {
		ctx.JSON(200, gin.H{
			"status_message": "Record Found'",
			"user":           user.EncodeJSON(),
		})
	}
}

func FindOnFile(ctx *gin.Context) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("controllers.users.FindOnFile()")
	params, serializeErr := serializeParam(ctx)
	fmt.Println("User param: ", params)
	fmt.Println("Error:?", serializeErr)
	if serializeErr != nil {
		fmt.Println("Malformed Register POST JSON parameter")
		sendErrResponse(serializeErr, ctx)
		return
	}

	user, findError := services.FindUserByEmailOnFile(params.User.Email)
	if findError != nil {
		sendErrResponse(findError, ctx)
	} else {
		ctx.JSON(200, gin.H{
			"status_message": "Record Found'",
			"user":           user.EncodeJSON(),
		})
	}
}
