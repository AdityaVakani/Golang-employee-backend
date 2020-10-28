package main

import (
	"employee_db/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/emp", controllers.GetAllEmp).Methods("GET")
	router.HandleFunc("/emp/{id}", controllers.GetEmp).Methods("GET")
	router.HandleFunc("/emp", controllers.CreateEmp).Methods("POST")
	router.HandleFunc("/emp/field/{id}", controllers.UpdateEmpField).Methods("PUT")
	router.HandleFunc("/emp/{id}", controllers.UpdateEmpAll).Methods("PUT")
	router.HandleFunc("/emp/{id}", controllers.DeleteEmp).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}
