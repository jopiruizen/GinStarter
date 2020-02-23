package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"encoding/json"
)

func BodyToJSONMap(ctx *gin.Context) (interface{}, error)  {
	body := ctx.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil  {
		fmt.Println(err.Error())
		return nil, err
	}

	var jsonMap interface{}
	json.Unmarshal(value,&jsonMap)
	fmt.Println("JSONBody: ", jsonMap)
	return jsonMap, err
}