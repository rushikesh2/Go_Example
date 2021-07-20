package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

type account struct {
	Name string `json:"name"`
}

// Create godoc
// @Summary Create a new employee
// @Description Create a new employee with the input paylod
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body account true "Create employee"
// @Success 200 {object} account
// @Router /user/account [get]
func Create(w http.ResponseWriter, r *http.Request) {
	log.Println("My test called")
	singleAccount := account{
		Name: "alibaba",
	}
	b, err := json.Marshal(singleAccount)
	if err != nil {
		fmt.Println("error:", http.StatusInternalServerError)
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}


// @title Orders API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /user/account
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/account", Create).Methods("GET")
	serverAddress := "localhost:5000"

	log.Println("starting server at", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
