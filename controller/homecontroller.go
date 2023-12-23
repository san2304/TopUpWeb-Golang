package controller

import (
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/navbar/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
