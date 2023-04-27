package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type server struct {
	db *sql.DB
}

type Manga struct {
	Id        int
	Name      string
	PageCount int
	Year      int
}

func database() server {
	database, _ := sql.Open("sqlite3", "anime.db")
	server := server{db: database}
	return server
}

func (s *server) seePage(w http.ResponseWriter, r *http.Request) {
	res, _ := s.db.Query("select * from manga;")
	var arr []Manga

	for res.Next() {
		var manga Manga
		if err := res.Scan(&manga.Id, &manga.Name, &manga.PageCount, &manga.Year); err != nil {
			log.Fatal(err)
		}
		arr = append(arr, manga)

	}

	tmpl, _ := template.ParseFiles("static/see.html")
	tmpl.Execute(w, arr)
}

func (s *server) addPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		pagecount := r.FormValue("pagecount")
		year := r.FormValue("year")
		_, err := s.db.Exec("insert into manga(name, pageCount, year) values ($1, $2, $3)", name, pagecount, year)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/see", http.StatusSeeOther)
	}
	tmpl, _ := template.ParseFiles("static/add.html")
	tmpl.Execute(w, nil)
}

func (s *server) deletePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		_, err := s.db.Exec("delete from manga where id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/see", http.StatusSeeOther)
	}
	tmpl, _ := template.ParseFiles("static/delete.html")
	tmpl.Execute(w, nil)
}

func (s *server) updatePage(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("name") == "" {
		id := r.FormValue("id")
		data := map[string]interface{}{"id": id}
		tmpl, _ := template.ParseFiles("static/update.html")
		tmpl.Execute(w, data)
		return
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	pagecount := r.FormValue("pagecount")
	year := r.FormValue("year")
	_, err := s.db.Exec("update manga set name=$1, pagecount=$2, year=$3 where id=$4;", name, pagecount, year, id)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/see", http.StatusSeeOther)
}

func main() {
	s := database()
	defer s.db.Close()

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/see", s.seePage)
	http.HandleFunc("/add", s.addPage)
	http.HandleFunc("/delete", s.deletePage)
	http.HandleFunc("/update", s.updatePage)

	http.ListenAndServe(":8080", nil)
}
