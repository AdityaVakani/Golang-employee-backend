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

func GetMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query(
		`with recursive cte as (
		  select     menuID,menuName,menuLink,menuComponent,menuVariant,menuIcon,
             parentID
		from       menu
		where      parentID = ? 
		union all
		select     m.menuID,m.menuName,m.menuLink,m.menuComponent,m.menuVariant,m.menuIcon,
					m.parentID
		from       menu m
		inner join cte
				on m.parentID = cte.menuID
	  )
	  select * from menu where menuID = ? union all select * from cte`, params["id"], params["id"])

	if err != nil {
		log.Println("error querying in GetMenu", err)
		return
	}

	defer result.Close()
	var allMenu []models.MenuItem

	for result.Next() {
		var menu models.MenuItem
		err := result.Scan(&menu.ID, &menu.Name, &menu.Link, &menu.Component, &menu.Variant, &menu.Icon, &menu.ParentID)
		if err != nil {
			log.Println("error reading select results", err)
			return
		}
		allMenu = append(allMenu, menu)

	}
	NewallMenu := models.MenuHirearchy(allMenu)
	//fmt.Println(NewallMenu)
	json.NewEncoder(w).Encode(NewallMenu)

}

func GetAllMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result1, err := db.Query("SELECT menuID FROM menu where parentID = 0")
	if err != nil {
		log.Println("Error selecting Menu ids", err)
	}

	var allMenus []models.MenuItem
	for result1.Next() {
		var row int
		err := result1.Scan(&row)
		if err != nil {
			log.Println("Error scanning main menu ids", err)
			return
		}
		result2, err := db.Query(`with recursive cte as (
			select     menuID,menuName,menuLink,menuComponent,menuVariant,menuIcon,
			   parentID
		  from       menu
		  where      parentID = ? 
		  union all
		  select     m.menuID,m.menuName,m.menuLink,m.menuComponent,m.menuVariant,m.menuIcon,
					  m.parentID
		  from       menu m
		  inner join cte
				  on m.parentID = cte.menuID
		)
		select * from menu where menuID = ? union all select * from cte`, row, row)

		var allMenu []models.MenuItem

		for result2.Next() {
			var menu models.MenuItem
			err := result2.Scan(&menu.ID, &menu.Name, &menu.Link, &menu.Component, &menu.Variant, &menu.Icon, &menu.ParentID)
			if err != nil {
				log.Println("error reading select results", err)
				return
			}
			allMenu = append(allMenu, menu)

		}
		NewallMenu := models.MenuHirearchy(allMenu)
		allMenus = append(allMenus, NewallMenu)

	}
	json.NewEncoder(w).Encode(allMenus)

}

func CreateMenu(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body", err)
		return
	}
	var menu models.MenuItem
	err = json.Unmarshal(body, &menu)
	if err != nil {
		log.Println("error unmarshaling", err)
		return
	}

	stmt, err := db.Prepare("INSERT INTO menu (menuID,menuName,menuLink,menuComponent,menuVariant,menuIcon,parentID) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Println("error preparing sql statement", err)
		return
	}
	_, err = stmt.Exec(menu.ID, menu.Name, menu.Link, menu.Component, menu.Variant, menu.Icon, menu.ParentID)

	if err != nil {
		log.Println("error executing sql statement", err)
		return
	}
	fmt.Fprintf(w, "New menu Created")

}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading update body", err)
		return
	}
	var menu models.MenuItem
	err = json.Unmarshal(body, &menu)
	if err != nil {
		log.Println("error unmarshaling update all json", err)
		return
	}
	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE menu SET menuID=?, menuName=?, menuLink=?, menuComponent=?, menuVariant=?, menuIcon=?, parentID=? where menuID =?")
	if err != nil {
		log.Println("error preparing update all statement", err)
		return
	}
	//fmt.Println(emp.FullName, emp.Email, emp.Mobile, emp.City, emp.Gender, emp.DepartmentID, emp.HireDate, emp.IsPermanent, params["id"])
	_, err = stmt.Exec(menu.ID, menu.Name, menu.Link, menu.Component, menu.Variant, menu.Icon, menu.ParentID, params["id"])
	if err != nil {
		log.Println("error executing update all statement", err)
		return
	}
	fmt.Fprintf(w, "Update all menu fields")
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare(`with recursive cte as (
		select     menuID,menuName,menuLink,menuComponent,menuVariant,menuIcon,
		   parentID
	  from       menu
	  where      parentID = ? 
	  union all
	  select     m.menuID,m.menuName,m.menuLink,m.menuComponent,m.menuVariant,m.menuIcon,
				  m.parentID
	  from       menu m
	  inner join cte
			  on m.parentID = cte.menuID
	)
	delete from menu where menuID in (select menuID from cte) or menuID = ?`)
	if err != nil {
		log.Println("error preparing delete statements", err)
		return
	}
	_, err = stmt.Exec(params["id"], params["id"])
	if err != nil {
		log.Println("error executing delete statements", err)
		return
	}
	fmt.Fprintf(w, "Menu deleted")

}
