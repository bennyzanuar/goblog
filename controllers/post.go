package controllers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetAllPost from service package
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	log.Info("route get")
}
