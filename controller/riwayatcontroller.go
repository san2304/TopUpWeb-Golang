package controller

import (
	"html/template"
	"net/http"
)

func Riwayat(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/navbar/riwayat.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
