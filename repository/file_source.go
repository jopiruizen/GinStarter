package repository

import (
	"encoding/json"
	"fmt"
	"go-restapi/models"
	"io/ioutil"
	"os"
)

type FileSource struct {
}

type jsonContent struct {
	/* note Unmarshal accept exportable properties */
	Users []models.User `json:"users"`
}

func (_ *FileSource) LoadSource() []models.User {

	fmt.Println("#############FileSource.LoadSource()")
	jsonFile, osErr := os.Open("data/users.json")
	if osErr != nil {
		fmt.Println("FileSource: Err:File Not Found", osErr)
	}

	jsonBytes, ioErr := ioutil.ReadAll(jsonFile)
	if ioErr != nil {
		fmt.Println("FileSource:Input Output Err", ioErr)
	}

	var content jsonContent
	json.Unmarshal(jsonBytes, &content)
	fmt.Println("Loaded Users: ", content.Users)
	return content.Users
}
