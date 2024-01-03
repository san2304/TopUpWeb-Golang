package controller

import (
	"html/template"
	"net/http"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/game/checkout.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
