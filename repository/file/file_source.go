package file

import (
	"encoding/json"
	log "github.com/golang/glog"
	"go-restapi/helper"
	"go-restapi/models"
	"io/ioutil"
	"os"
)

type FileSource struct {
	usersList []models.User
}

type jsonContent struct {
	/* note Unmarshal accept exportable properties */
	Users []models.User `json:"users"`
}

func (src *FileSource) LoadSource() {
	log.Info("FileSource.LoadSource()")
	log.Info("#############FileSource.LoadSource()")
	jsonFile, osErr := os.Open("data/users.json")
	if osErr != nil {
		log.Error("FileSource: Err:File Not Found", osErr.Error())
	}

	jsonBytes, ioErr := ioutil.ReadAll(jsonFile)
	if ioErr != nil {
		log.Error("FileSource:Input Output Err", ioErr.Error())
	}

	var content jsonContent
	json.Unmarshal(jsonBytes, &content)
	log.Info("Loaded Users: ", content.Users)
	src.usersList = content.Users
}

func (src *FileSource) Find(email string) (models.User, error) {
	return helper.SearchByEmail(email, src.usersList)
}
