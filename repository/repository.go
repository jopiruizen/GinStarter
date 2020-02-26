package repository

import (
	log "github.com/golang/glog"
	"go-restapi/models"
	"go-restapi/repository/file"
	"go-restapi/repository/static"
)

type _sourceType struct {
	STATIC string
	FILE   string
}

var SourceType = _sourceType{
	STATIC: "static",
	FILE:   "file",
}

type IRepoSource interface {
	LoadSource()
	Find(email string) (models.User, error)
}

func NewRepoSource(srcType string) IRepoSource {
	var source IRepoSource
	log.Info("NewRepoSource()", srcType)
	switch srcType {

	case SourceType.STATIC:
		source = &static.StaticSource{}

	case SourceType.FILE:
		source = &file.FileSource{}

	}
	source.LoadSource()
	return source
}
