package users

import (
	"github.com/gin-gonic/gin"
	"go-restapi/models"
)

func sendErrResponse(err error, ctx *gin.Context) {
	switch err {
	case models.ErrBadInput:
		ctx.JSON(404, gin.H{
			"status_message": "Bad Input Error",
			"error":          err.Error(),
		})

	case models.ErrNoRecordFound:
		ctx.JSON(422, gin.H{
			"status_message": "No Record Found",
			"error":          err.Error(),
		})
	default:
		ctx.JSON(422, gin.H{
			"status_message": "Unknown Server/Technical Error",
			"error":          err.Error(),
		})
	}

}
