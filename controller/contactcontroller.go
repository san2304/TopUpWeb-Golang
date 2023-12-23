package controller

import (
	"html/template"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/navbar/contact.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
