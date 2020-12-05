package api

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/viewuser.html","templates/edituser.html"))


var TempUser = User{Name: "Deepak Singhvi", Email: "deepak.singhvi@gmail.com"}
func HandleRequest() {
	log.Println("Web App Started")
	http.HandleFunc("/", viewUserHandler)
	http.HandleFunc("/viewuser", viewUserHandler)
	http.HandleFunc("/userform", editUserFormHandler)
	http.HandleFunc("/updateuser", updateUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewUserHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	renderTemplate(w,"viewuser",TempUser)
}


func editUserFormHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	renderTemplate(w,"edituser",TempUser)
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// form is not sending Content-Type as json
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	TempUser.Name = r.Form.Get("username")
	TempUser.Email = r.Form.Get("emailid")
	log.Println(TempUser)
	renderTemplate(w,"viewuser",TempUser)
}


