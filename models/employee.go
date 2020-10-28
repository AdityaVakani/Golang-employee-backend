package models

import (
	"employee_db/connectdb"
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var db = connectdb.DB
var err error

type EmpData struct {
	ID           int    `json:"id"`
	FullName     string `json:"fullName" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Mobile       string `json:"mobile" validate:"required,min=9,max=15"`
	City         string `json:"city" validate:"required"`
	Gender       string `json:"gender" validate:"required"`
	DepartmentID string `json:"departmentId" validate:"required,validDep"`
	HireDate     string `json:"hireDate" validate:"required"`
	IsPermanent  bool   `json:"isPermanent"`
}

func validDep(dep validator.FieldLevel) bool {
	result, err := db.Query("SELECT exists(SELECT* FROM department_lu where depID=?)", dep.Field().String())
	if err != nil {
		log.Println("error querying department_lu", err)
		return false
	}
	var res int
	for result.Next() {
		err = result.Scan(&res)
		if err != nil {
			log.Println("error reading depid results")
			return false
		}
	}
	return res == 1

}
func Validation(emp EmpData) map[string]string {
	//fmt.Println(emp)
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Println("translator not found")
	}
	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Println(err)
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
