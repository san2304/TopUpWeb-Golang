package controller

import (
	"html/template"
	"net/http"
)

func Invoice(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/navbar/invoice.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
