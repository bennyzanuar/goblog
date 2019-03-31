package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bennyzanuar/goblog/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()
	routes.InitPostRoutes(v1)

	port := os.Getenv("app_port")
	if port == "" {
		port = "8004"
	}

	err := http.ListenAndServe(":"+port, appendTrailingSlash(r))
	if err != nil {
		log.Panic("Error", err)
	}
	fmt.Println("server running")
}

func appendTrailingSlash(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
