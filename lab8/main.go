package main

import (
	"html/template"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, nil)
}

func watchPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/watch.html")
	t.Execute(w, nil)
}

func buyPage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			data = map[string]interface{}{"error": err}
			t, _ := template.ParseFiles("static/buy.html")
			t.Execute(w, data)
			return
		}

		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		data = map[string]interface{}{"first": firstName, "last": lastName}
		t, _ := template.ParseFiles("static/buyComplete.html")
		t.Execute(w, data)
		return
	}

	t, _ := template.ParseFiles("static/buy.html")
	t.Execute(w, nil)
}

func buyCompletePage(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/watch", watchPage)
	http.HandleFunc("/buy", buyPage)
	http.HandleFunc("/buyComplete", buyCompletePage)

	http.ListenAndServe(":8000", nil)
}
