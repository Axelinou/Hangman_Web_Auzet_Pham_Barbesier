package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Email     string
	LastName  string
	FirstName string
	Success   bool
}

func main() {
	tmpl1 := template.Must(template.ParseFiles("index2.html"))
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
		details := User{
			Email:     r.FormValue("email"),
			LastName:  r.FormValue("lastname"),
			FirstName: r.FormValue("firstname"),
			Success:   true,
		}
		tmpl1.Execute(w, details)
	})
	http.ListenAndServe(":80", nil)

}