package services

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql" // sql driver
)

//player is the model of player
type player struct {
	Id       int
	Name     string
	LastName string
	Number   string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "9990"
	dbName := "team"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

//Index template
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM players ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	// p is a player struct
	p := player{}
	res := []player{}
	for selDB.Next() {
		var id int
		var name, lastName, number string
		err = selDB.Scan(&id, &name, &lastName, &number)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.LastName = lastName
		p.Number = number
		res = append(res, p)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//Show template
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM players WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	p := player{}
	for selDB.Next() {
		var id int
		var name, lastName, number string
		err = selDB.Scan(&id, &name, &lastName, &number)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.LastName = lastName
	}
	tmpl.ExecuteTemplate(w, "Show", p)
	defer db.Close()
}

//New players
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//Edit players
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM players WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	p := player{}
	for selDB.Next() {
		var id int
		var name, lastName, number string
		err = selDB.Scan(&id, &name, &lastName, &number)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.LastName = lastName
		p.Number = number
	}
	tmpl.ExecuteTemplate(w, "Edit", p)
	defer db.Close()
}

//Insert players
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		lastName := r.FormValue("lastName")
		number := r.FormValue("number")
		insForm, err := db.Prepare("INSERT INTO players(name, lastName, number) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, lastName, number)
		log.Println("INSERT: Name: " + name + " | Last Name: " + lastName + " | Number: " + number)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//Update players
func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		lastName := r.FormValue("lastName")
		number := r.FormValue("number")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE players SET name=?, lastName=?, number=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, lastName, number, id)
		log.Println("UPDATE: Name: " + name + " | Last Name: " + lastName + " | Number: " + number)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//Delete players
func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	p := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM players WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(p)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
