package main

import (
	log "github.com/golang/glog"
	"go-restapi/routes"
)

func main() {
	log.Info("Initializing Gin Services...")
	router := routes.SetupRouter()
	router.Run()
}
