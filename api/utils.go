package api

import "net/http"

type User struct {
	Name string `json:"username"`
	Email string `json:"emailid"`
}


func renderTemplate(w http.ResponseWriter, tmpl string, user User) {
	err := templates.ExecuteTemplate(w, tmpl+".html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}