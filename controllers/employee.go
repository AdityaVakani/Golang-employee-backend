package controllers

import (
	"employee_db/connectdb"
	"employee_db/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

var db = connectdb.DB
var err error

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "INDEX Found")
}

func GetAllEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT * FROM employee_db")
	if err != nil {
		log.Println("error selecting all from table", err)
		return
	}
	defer result.Close()
	var allEmp []models.EmpData

	for result.Next() {
		var emp models.EmpData
		err := result.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Mobile, &emp.City, &emp.Gender, &emp.DepartmentID, &emp.HireDate, &emp.IsPermanent)
		if err != nil {
			log.Println("error reading select results", err)
			return
		}
		allEmp = append(allEmp, emp)

	}

	json.NewEncoder(w).Encode(allEmp)

}

func GetEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM employee_db WHERE empID = ?", params["id"])
	if err != nil {
		log.Println("error selecting from table", err)
		return
	}
	defer result.Close()
	var emp models.EmpData

	for result.Next() {
		err := result.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Mobile, &emp.City, &emp.Gender, &emp.DepartmentID, &emp.HireDate, &emp.IsPermanent)
		if err != nil {
			log.Println("error reading select results", err)
			return
		}
	}

	json.NewEncoder(w).Encode(emp)

}

func CreateEmp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body", err)
		return
	}
	var emp models.EmpData
	err = json.Unmarshal(body, &emp)
	if err != nil {
		log.Println("error unmarshaling", err)
		return
	}
	fmt.Println(reflect.TypeOf(emp.HireDate), emp)
	errorMap := models.Validation(emp)
	if len(errorMap) != 0 {
		fmt.Fprint(w, errorMap)
	} else {
		stmt, err := db.Prepare("INSERT INTO employee_db (empName,empEmail,empPhone,empCity,empGender,empDepartmentID,empHireDate,empIsPermanent) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			log.Println("error preparing sql statement", err)
			return
		}
		_, err = stmt.Exec(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent)

		if err != nil {
			log.Println("error executing sql statement", err)
			return
		}
		fmt.Fprintf(w, "New Employee Created")

	}
}

func UpdateEmpField(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body in update", err)
		return
	}
	keyVal := make(map[string]string)
	err = json.Unmarshal(body, &keyVal)
	if err != nil {
		log.Println("error unmarshaling body in update", err)
		return
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
		log.Println("error preparing update sql", err)
		return
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("error executing update statement", err)
		return
	}
	fmt.Fprintf(w, "Employee was updated")
}

func UpdateEmpAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading update all body", err)
		return
	}
	var emp models.EmpData
	err = json.Unmarshal(body, &emp)
	if err != nil {
		log.Println("error unmarshaling update all json", err)
		return
	}
	params := mux.Vars(r)
	errorMap := models.Validation(emp)
	if len(errorMap) != 0 {
		fmt.Fprint(w, errorMap)
	} else {
		stmt, err := db.Prepare("UPDATE employee_db SET empName=?, empEmail=?, empPhone=?, empCity=?, empGender=?, empDepartmentID=?, empHireDate=?, empIsPermanent=? where empID =?")
		if err != nil {
			log.Println("error preparing update all statement", err)
			return
		}
		//fmt.Println(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent, params["id"])
		_, err = stmt.Exec(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent, params["id"])
		if err != nil {
			log.Println("error executing update all statement", err)
			return
		}
		fmt.Fprintf(w, "Update all employee fields")
	}

}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM employee_db WHERE empID = ?")
	if err != nil {
		log.Println("error preparing delete statements", err)
		return
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		log.Println("error executing delete statements", err)
		return
	}
	fmt.Fprintf(w, "employee deleted")

}
