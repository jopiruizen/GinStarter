package main

import "fmt"
import "github.com/gin-gonic/gin"
import "go-restapi/controllers/users"

func homePage (context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "This is a home page generated by get.",
	})
}

func main() {
	fmt.Println("Initializing Gin Services...")
	r := gin.Default()
	r.GET("/", homePage)
	r.POST("/user/register", users.Register)
	r.Run()
}