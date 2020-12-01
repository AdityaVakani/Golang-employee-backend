package controllers

import (
	"employee_db/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Println("error selecting all from orders table", err)
		return
	}
	defer result.Close()

	var allOrders []models.Orders

	for result.Next() {
		var order models.Orders
		err := result.Scan(&order.OrderID, &order.Customer, &order.Address, &order.City, &order.Gender, &order.OrderDate, &order.IsDelivered)
		if err != nil {
			log.Println("error reading select results", err)
			return
		}
		allOrders = append(allOrders, order)

	}

	json.NewEncoder(w).Encode(allOrders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM orders WHERE orderID = ?", params["orderId"])
	if err != nil {
		log.Println("error selecting orders from table", err)
		return
	}
	defer result.Close()
	var order models.Orders

	for result.Next() {
		err := result.Scan(&order.OrderID, &order.Customer, &order.Address, &order.City, &order.Gender, &order.OrderDate, &order.IsDelivered)
		if err != nil {
			log.Println("error reading scanning results", err)
			return
		}
	}

	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body", err)
		return
	}
	var order models.Orders
	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Println("error unmarshaling", err)
		return
	}

	stmt, err := db.Prepare("INSERT INTO orders (customer,address,city,gender,orderDate,isDelivered) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Println("error preparing sql statement", err)
		return
	}
	_, err = stmt.Exec(order.Customer, order.Address, order.City, order.Gender, order.OrderDate, order.IsDelivered)

	if err != nil {
		log.Println("error executing sql statement", err)
		return
	}
	fmt.Fprintf(w, "New Order Created")

}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading update all body", err)
		return
	}
	var order models.Orders
	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Println("error unmarshaling update  json", err)
		return
	}
	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE orders SET customer=?,address=?,city=?,gender=?,orderDate=?,isDelivered=? where orderID =?")
	if err != nil {
		log.Println("error preparing update all statement", err)
		return
	}

	_, err = stmt.Exec(order.Customer, order.Address, order.City, order.Gender, order.OrderDate, order.IsDelivered, params["orderId"])
	if err != nil {
		log.Println("error executing update all statement", err)
		return
	}
	fmt.Fprintf(w, "Updated order")

}

func DeleteOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM orders,order_details using orders INNER JOIN order_details ON orders.orderID = order_details.orderID WHERE orders.orderID = ?")
	if err != nil {
		log.Println("error preparing delete statements", err)
		return
	}
	_, err = stmt.Exec(params["orderId"])
	if err != nil {
		log.Println("error executing delete statements", err)
		return
	}
	fmt.Fprintf(w, "order deleted")
}

func GetAllDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM order_details WHERE orderID =?", params["orderId"])
	if err != nil {
		log.Println("error selecting all order details from table", err)
		return
	}
	defer result.Close()

	var allDetails []models.OrderDetails

	for result.Next() {
		var row models.OrderDetails
		err := result.Scan(&row.ProductID, &row.OrderID, &row.Product, &row.Quantity, &row.Price)
		if err != nil {
			log.Println("error reading select results", err)
			return
		}
		allDetails = append(allDetails, row)

	}

	json.NewEncoder(w).Encode(allDetails)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM order_details WHERE orderID =? AND productID =?", params["orderId"], params["productId"])
	if err != nil {
		log.Println("error selecting order details from table", err)
		return
	}
	defer result.Close()

	var details models.OrderDetails

	for result.Next() {
		err := result.Scan(&details.ProductID, &details.OrderID, &details.Product, &details.Quantity, &details.Price)
		if err != nil {
			log.Println("error reading select results", err)
			return
		}

	}

	json.NewEncoder(w).Encode(details)
}

func AddDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body", err)
		return
	}
	var detail models.OrderDetails
	err = json.Unmarshal(body, &detail)
	if err != nil {
		log.Println("error unmarshaling", err)
		return
	}
	params := mux.Vars(r)
	stmt, err := db.Prepare("INSERT INTO order_details (orderID,product,quantity,price) VALUES(?,?,?,?)")
	if err != nil {
		log.Println("error preparing sql statement", err)
		return
	}
	_, err = stmt.Exec(params["orderId"], detail.Product, detail.Quantity, detail.Price)

	if err != nil {
		log.Println("error executing sql statement", err)
		return
	}
	fmt.Fprintf(w, "New Detail Added")
}

func UpdateDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading update body", err)
		return
	}
	var detail models.OrderDetails
	err = json.Unmarshal(body, &detail)
	if err != nil {
		log.Println("error unmarshaling update  json", err)
		return
	}
	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE order_details SET product=?,quantity=?,price=? where orderID =? and productID=?")
	if err != nil {
		log.Println("error preparing update statement", err)
		return
	}

	_, err = stmt.Exec(detail.Product, detail.Quantity, detail.Price, params["orderId"], params["productId"])
	if err != nil {
		log.Println("error executing update all statement", err)
		return
	}
	fmt.Fprintf(w, "Updated details")

}
func DeleteDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM order_details WHERE orderID = ? and productID=?")
	if err != nil {
		log.Println("error preparing delete statements", err)
		return
	}
	_, err = stmt.Exec(params["orderId"], params["productId"])
	if err != nil {
		log.Println("error executing delete statements", err)
		return
	}
	fmt.Fprintf(w, "detail deleted")

}
