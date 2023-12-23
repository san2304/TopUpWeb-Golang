package controller

import (
	"html/template"
	"net/http"
	"topup-game/entities"
	"topup-game/models"
)

func Freefire(w http.ResponseWriter, r *http.Request) {
	diamondOptions, err := models.GetDiamondOptions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("views/game/ff.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, struct{ DiamondOptions []entities.DiamondOption }{DiamondOptions: diamondOptions})
}
