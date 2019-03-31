package controllers

import (
	"fmt"
	"net/http"

	services "github.com/bennyzanuar/goblog/implements"
	"github.com/sanity-io/litter"
)

// GetAllPost from service package
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	data := services.FindAll()
	fmt.Printf("%+v", data)
	litter.Dump(data)
}
