package users

import (
	"github.com/gin-gonic/gin"
	"go-restapi/models"
)

func sendErrResponse(err error, ctx *gin.Context) {
	switch err {
	case models.ErrBadInput:
		ctx.JSON(200, gin.H{
			"status_message": "Bad Input error ",
			"error":          err,
		})

	case models.ErrNoRecordFound:
		ctx.JSON(200, gin.H{
			"status_message": "No Record Found",
			"error":          err,
		})
	}

}
