package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

//EmpData is Employee data
type EmpData struct {
	ID           int    `json:"id"`
	FullName     string `json:"fullName" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Mobile       string `json:"mobile" validate:"required,min=9,max=15"`
	City         string `json:"city" validate:"required"`
	Gender       string `json:"gender" validate:"required"`
	DepartmentID int    `json:"departmentId" validate:"required,validDep"`
	HireDate     string `json:"hireDate" validate:"required"`
	IsPermanent  bool   `json:"isPermanent"`
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
		log.Fatalln("error pinging db", err)
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/emp", getAllEmp).Methods("GET")
	router.HandleFunc("/emp/{id}", getEmp).Methods("GET")
	router.HandleFunc("/emp", createEmp).Methods("POST")
	router.HandleFunc("/emp/field/{id}", updateEmpField).Methods("PUT")
	router.HandleFunc("/emp/{id}", updateEmpAll).Methods("PUT")
	router.HandleFunc("/emp/{id}", deleteEmp).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "INDEX Found")
}

func getAllEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT * FROM employee_db")
	if err != nil {
		log.Fatalln("error selecting all from table", err)
	}
	defer result.Close()
	var allEmp []EmpData

	for result.Next() {
		var emp EmpData
		err := result.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Mobile, &emp.City, &emp.Gender, &emp.DepartmentID, &emp.HireDate, &emp.IsPermanent)
		if err != nil {
			log.Fatalln("error reading select results", err)
		}
		allEmp = append(allEmp, emp)

	}

	json.NewEncoder(w).Encode(allEmp)

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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("error reading body", err)
	}
	var emp EmpData
	err = json.Unmarshal(body, &emp)
	if err != nil {
		log.Fatalln("error unmarshaling", err)
	}
	fmt.Println(reflect.TypeOf(emp.HireDate), emp)
	errorMap := validation(emp)
	if len(errorMap) != 0 {
		fmt.Fprint(w, errorMap)
	} else {
		stmt, err := db.Prepare("INSERT INTO employee_db (empName,empEmail,empPhone,empCity,empGender,empDepartmentID,empHireDate,empIsPermanent) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			log.Fatalln("error preparing sql statement", err)
		}
		_, err = stmt.Exec(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent)

		if err != nil {
			log.Fatalln("error executing sql statement", err)
		}
		fmt.Fprintf(w, "New Employee Created")

	}
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
	errorMap := validation(emp)
	if len(errorMap) != 0 {
		fmt.Fprint(w, errorMap)
	} else {
		stmt, err := db.Prepare("UPDATE employee_db SET empName=?, empEmail=?, empPhone=?, empCity=?, empGender=?, empDepartmentID=?, empHireDate=?, empIsPermanent=? where empID =?")
		if err != nil {
			log.Fatalln("error preparing update all statement", err)
		}
		//fmt.Println(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent, params["id"])
		_, err = stmt.Exec(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent, params["id"])
		if err != nil {
			log.Fatalln("error executing update all statement", err)
		}
		fmt.Fprintf(w, "Update all employee fields")
	}

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
func validDep(dep validator.FieldLevel) bool {
	result, err := db.Query("SELECT exists(SELECT* FROM department_lu where depID=?)", dep.Field().Int())
	if err != nil {
		log.Fatalln("error querying department_lu", err)
	}
	var res int
	for result.Next() {
		err = result.Scan(&res)
		if err != nil {
			log.Fatalln("error reading depid results")
		}
	}
	return res == 1

}
func validation(emp EmpData) map[string]string {
	//fmt.Println(emp)
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatalln("translator not found")
	}
	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatalln(err)
	}
	_ = v.RegisterTranslation("validDep", trans, func(ut ut.Translator) error {
		return ut.Add("validDep", "{0} must be a valid department", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("validDep", fe.Field())
		return t
	})
	_ = v.RegisterValidation("validDep", validDep)
	err = v.Struct(emp)
	errorMap := make(map[string]string)
	if err != nil {
		errorMap = err.(validator.ValidationErrors).Translate(trans)
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e.Translate(trans))
		}
	} else {
		fmt.Println("valid input")
	}
	return errorMap

}
