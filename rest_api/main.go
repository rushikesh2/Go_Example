package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest_api/config"
	"rest_api/dao"
	"rest_api/model"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/gorilla/mux"
)

var confi = config.Config{}
var da = dao.MongoDao{}

// Create godoc
// @Summary Create a new employee
// @Description Create a new employee with the input paylod
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body model.Employee true "Create employee"
// @Success 200 {object} model.Employee
// @Router /user/account [post]
func Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Employee register api called")

	var user model.Employee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Account unable to decipher %v", err), http.StatusBadRequest)
		log.Print("Invalid Account unable to decipher", http.StatusBadRequest)
	}

	err = da.Insert(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while insert employee into collection %v", err), http.StatusInternalServerError)
		log.Print("Unable to insert record into database")
		return
	}
	usr, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usr)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var user model.Employee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Account unable to decipher %v", err), http.StatusBadRequest)
		log.Print("Invalid Account unable to decipher", http.StatusBadRequest)
	}
	err = da.UpdateEmp(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while insert employee into collection: %s,  %v", dao.COLLECTION, err), http.StatusInternalServerError)
		log.Print("Unable to insert record into database")
		return
	}
	json.NewEncoder(w).Encode(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Delete godoc
// @Summary Delete a employee
// @Description delete employee by ID
// @ID get-string-by-int
// @Produce  json
// @Param id path int true "empid"
// @Success 200 {string} body
// @Router /user/account/{empid} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	empid := vars["empid"]

	err := da.Delete(empid)
	if err != nil {
		http.Error(w, fmt.Sprintf("DeleteAccount :: error occurred, Account :: %s", empid), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body := "{\"message\":\"Delete request successfully submitted for processing\"}"
	_, err = w.Write([]byte(body))
	if err != nil {
		panic(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid := vars["empid"]
	decode := json.NewDecoder(r.Body)
	var user_update map[string]interface{}

	empdetails, err := da.FindById(empid)
	if err != nil {
		http.Error(w, fmt.Sprintf("FindRecored :: error occurred, Employee :: %s", empid), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%s", empdetails)

	err = decode.Decode(&user_update)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Account unable to decipher %v", err), http.StatusBadRequest)
		log.Print("Invalid Account unable to decipher", http.StatusBadRequest)
	}

	err = checkUpdate(&empdetails, user_update)
	if err != nil {
		log.Printf("Error while checking update record  for empployee: %s \n Error: %v", empdetails.EmpID, err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	fmt.Printf("%s", empdetails)
	err = da.Update(empdetails)
	if err != nil {
		http.Error(w, fmt.Sprintf("UpdateAccount :: error occurred, Employee :: %s", empid), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(empdetails)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func checkUpdate(emp *model.Employee, user_update map[string]interface{}) (err error) {
	if _, ok := user_update["mail"]; ok {
		emp.Email, ok = user_update["mail"].(string)
		log.Panicln("not found", ok)
	}
	if _, ok := user_update["name"]; ok {
		emp.Name, ok = user_update["name"].(string)
		log.Panicln("not found", ok)
	}
	if _, ok := user_update["position"]; ok {
		emp.Position, ok = user_update["position"].(string)
		log.Panicln("not found", ok)
	}
	if _, ok := user_update["phone"]; ok {
		emp.Phone, ok = user_update["phone"].(string)
		log.Panicln("not found", ok)
	}
	if _, ok := user_update["practice"]; ok {
		emp.Practice, ok = user_update["practice"].(string)
		log.Panicln("not found", ok)
	}
	if _, ok := user_update["empid"]; ok {
		emp.EmpID, ok = user_update["empid"].(string)
		log.Panicln("not found", ok)
	}
	return
}

// Read godoc
// @Summary Show a account
// @Description get employee by empid
// @ID get-string-by-int
// @Produce  json
// @Param id path int true "empid"
// @Success 200 {object} model.Employee
// @Router /user/account/{empid} [get]
func Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid := vars["empid"]

	usr, err := da.FindById(empid)
	if err != nil {
		http.Error(w, "Error while querying database for reading record:", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(usr)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// ReadAll godoc
// @Summary Show all account
// @Description get all employees
// @Produce  json
// @Success 200 {array} model.Employee
// @Router /user/account [get]
func ReadAll(w http.ResponseWriter, r *http.Request) {
	var usr []model.Employee

	usr, err := da.FindAll()
	if err != nil {
		http.Error(w, "Error while querying database for reading all record:", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(usr)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	confi.Read()

	da.Server = confi.Server
	da.Database = confi.Database
	da.Connect()
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

	r.HandleFunc("/user/account", Create).Methods("POST")
	r.HandleFunc("/user/account", ReadAll).Methods("GET")
	r.HandleFunc("/user/account", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/user/account/{empid}", Delete).Methods("DELETE")
	r.HandleFunc("/user/account/{empid}", Update).Methods("PATCH")
	r.HandleFunc("/user/account/{empid}", Read).Methods("GET")

	serverAddress := "localhost:5000"

	// srv := &http.Server{
	// 	Handler:      r,
	// 	Addr:         "127.0.0.1:5000",
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	log.Println("starting server at", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//log.Fatal(srv.ListenAndServe())

}
