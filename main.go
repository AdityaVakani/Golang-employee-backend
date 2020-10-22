package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//EmpData is Employee data
type EmpData struct {
	ID           int    `json:"Id"`
	FullName     string `json:"FullName"`
	Email        string `json:"Email"`
	Mobile       int    `json:"Mobile"`
	City         string `json:"City"`
	Gender       string `json:"Gender"`
	DepartmentID int    `json:"DepartmentID"`
	HireDate     string `json:"HireDate"`
	IsPermanent  int    `json:"IsPermanent"`
}

var db *sql.DB
var err error

func init() {
	datasource := "root:password@tcp(localhost:3306)/go_backend01"
	db, err = sql.Open("mysql", datasource)
	if err != nil {
		log.Fatalln("error connecting to db", err)
	}
	fmt.Println("db is connected")
	err = db.Ping()
	if err != nil {
		log.Fatalln("error pining db", err)
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/emp/{id}", getEmp).Methods("GET")
	router.HandleFunc("/emp", createEmp).Methods("POST")
	router.HandleFunc("/emp/{id}", updateEmpField).Methods("PUT")
	router.HandleFunc("/emp/update-all/{id}", updateEmpAll).Methods("PUT")
	router.HandleFunc("/emp/{id}", deleteEmp).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "INDEX Found")
}

func getEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM employee_db WHERE empID = ?", params["id"])
	if err != nil {
		log.Fatalln("error selecting from table", err)
	}
	defer result.Close()
	var emp EmpData

	for result.Next() {
		err := result.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Mobile, &emp.City, &emp.Gender, &emp.DepartmentID, &emp.HireDate, &emp.IsPermanent)
		if err != nil {
			log.Fatalln("error reading select results", err)
		}
	}

	json.NewEncoder(w).Encode(emp)

}

func createEmp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO employee_db (empName,empEmail,empPhone,empCity,empGender,empDepartmentID,empHireDate,empIsPermanent) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatalln("error preparing sql statement", err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("error reading body", err)
	}
	var emp EmpData
	err = json.Unmarshal(body, &emp)
	if err != nil {
		log.Fatalln("error unmarshaling", err)
	}
	_, err = stmt.Exec(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent)

	if err != nil {
		log.Fatalln("error executing sql statement", err)
	}
	fmt.Fprintf(w, "New Employee Created")

}

func updateEmpField(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("error reading body in update", err)
	}
	keyVal := make(map[string]string)
	err = json.Unmarshal(body, &keyVal)
	if err != nil {
		log.Fatalln("error unmarshaling body in update", err)
	}
	colName := keyVal["columnName"]
	colValue := keyVal["columnValue"]
	if reflect.TypeOf(colValue) == reflect.TypeOf("Str") {
		colValue = "'" + colValue + "'"
	}
	sqlStatement := "UPDATE employee_db SET " + string(colName) + " = " + string(colValue) + " where empID = " + string(params["id"])
	//fmt.Println(sqlStatement)
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		log.Fatalln("error preparing update sql", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("error executing update statement", err)
	}
	fmt.Fprintf(w, "Employee was updated")
}

func updateEmpAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("UPDATE employee_db SET empName=?, empEmail=?, empPhone=?, empCity=?, empGender=?, empDepartmentID=?, empHireDate=?, empIsPermanent=? where empID =?")
	if err != nil {
		log.Fatalln("error preparing update all statement", err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("error reading update all body", err)
	}
	var emp EmpData
	err = json.Unmarshal(body, &emp)
	if err != nil {
		log.Fatalln("error unmarshaling update all json", err)
	}
	params := mux.Vars(r)
	fmt.Println(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent, params["id"])
	_, err = stmt.Exec(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent, params["id"])
	if err != nil {
		log.Fatalln("error executing update all statement", err)
	}
	fmt.Fprintf(w, "Update all employee")
}

func deleteEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM employee_db WHERE empID = ?")
	if err != nil {
		log.Fatalln("error preparing delete statements", err)
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "employee deleted")

}
