package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-restapi/services"
)

func Find(ctx *gin.Context) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("controllers.users.Find()")
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
