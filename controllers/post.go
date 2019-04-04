package controllers

import (
	"net/http"

	srv "github.com/bennyzanuar/goblog/services"
	u "github.com/bennyzanuar/goblog/utils"
)

// GetAllPost from service package
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	data := srv.FindAllPost(1, 10)
	u.Respond(w, data.Result, data.ParamsOutput)
}
