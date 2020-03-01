package repository

import (
	log "github.com/golang/glog"
	"go-restapi/models"
	"go-restapi/repository/commons"
	"go-restapi/repository/file"
	"go-restapi/repository/static"
)

type IRepoSource interface {
	LoadSource()
	Find(email string) (models.User, error)
}

func NewRepoSource(srcType string) IRepoSource {
	var source IRepoSource
	log.Info("NewRepoSource()", srcType)
	switch srcType {

	case commons.SOURCE_TYPE_STATIC:
		source = &static.StaticSource{}

	case commons.SOURCE_TYPE_FILE:
		source = &file.FileSource{}

	}
	source.LoadSource()
	return source
}
