package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type server struct {
	db *sql.DB
}

func dbConnect() server {
	db, err := sql.Open("sqlite3", "library.db")
	fmt.Println("Opening database")
	if err != nil {
		log.Fatal(err)
	}

	s := server{db: db}

	return s
}

func (s *server) add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		pageCount := r.FormValue("pages")
		year := r.FormValue("year")
		_, err := s.db.Exec("insert into manga(name, pageCount, year) values ($1, $2, $3)", name, pageCount, year)
		if err != nil {
			log.Fatal(err)
		}
		data := map[string]interface{}{"name": name, "text": " added!"}
		t, _ := template.ParseFiles("static/add.html")
		t.Execute(w, data)
		return
	}
	t, _ := template.ParseFiles("static/add.html")
	t.Execute(w, nil)
}

func main() {
	s := dbConnect()
	defer s.db.Close()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/add", s.add)

	http.ListenAndServe(":8080", nil)
}
