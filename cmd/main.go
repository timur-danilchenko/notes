package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timur-danilchenko/go-nodes/api"
)

func main() {
	router := mux.NewRouter()
	api.RegisterRoutes(router)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
