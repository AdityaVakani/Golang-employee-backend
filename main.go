package main

import (
	"employee_db/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/employee", controllers.GetAllEmp).Methods("GET")
	router.HandleFunc("/employee/{id}", controllers.GetEmp).Methods("GET")
	router.HandleFunc("/employee", controllers.CreateEmp).Methods("POST")
	router.HandleFunc("/employee/field/{id}", controllers.UpdateEmpField).Methods("PUT")
	router.HandleFunc("/employee/{id}", controllers.UpdateEmpAll).Methods("PUT")
	router.HandleFunc("/employee/{id}", controllers.DeleteEmp).Methods("DELETE")

	router.HandleFunc("/menu_row/{id}", controllers.GetMenuRow).Methods("GET")
	router.HandleFunc("/menu/{id}", controllers.GetMenu).Methods("GET")
	router.HandleFunc("/menu", controllers.GetAllMenu).Methods("GET")
	router.HandleFunc("/menu_tree/{id}", controllers.GetMenuTree).Methods("GET")
	router.HandleFunc("/menu_tree", controllers.GetAllMenuTree).Methods("GET")
	router.HandleFunc("/menu", controllers.CreateMenu).Methods("POST")
	router.HandleFunc("/menu/{id}", controllers.UpdateMenu).Methods("PUT")
	router.HandleFunc("/menu/{id}", controllers.DeleteMenu).Methods("DELETE")

	router.HandleFunc("/orders", controllers.GetAllOrders).Methods("GET")
	router.HandleFunc("/orders/{orderId}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{orderId}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", controllers.DeleteOrders).Methods("DELETE")

	router.HandleFunc("/orders/{orderId}/order_details", controllers.GetAllDetails).Methods("GET")
	router.HandleFunc("/orders/{orderId}/order_details/{productId}", controllers.GetDetail).Methods("GET")
	router.HandleFunc("/orders/{orderId}/order_details", controllers.AddDetails).Methods("POST")
	router.HandleFunc("/orders/{orderId}/order_details/{productId}", controllers.UpdateDetails).Methods("PUT")
	router.HandleFunc("/orders/{orderId}/order_details/{productId}", controllers.DeleteDetails).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
