package gamecontroller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"topup-game/entities"
	"topup-game/models"
)

func Mlbb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		product, err := models.GetProductMlbb()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("views/game/mlbb.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct{ Products []entities.InvoiceWithProduct }{Products: product})
	}

	if r.Method == "POST" {
		var invoice entities.InvoiceWithProduct

		dataID, err := strconv.Atoi(r.FormValue("dataID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		productID, err := strconv.Atoi(r.FormValue("product_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		phoneNumber, err := strconv.ParseInt(r.FormValue("phone_number"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		invoice.DataID = dataID
		invoice.ProductID = productID
		invoice.Payment = r.FormValue("payment")
		invoice.PhoneNumber = phoneNumber
		invoice.Status = "Belum Dibayar"
		invoice.Time = time.Now()

		if ok := models.CreateInvoice(&invoice); !ok {
			log.Println("Error: Unable to create invoice")
			http.Error(w, "Unable to create invoice", http.StatusInternalServerError)
			return
		}

		log.Println("Debug: Invoice created successfully")
		http.Redirect(w, r, "/checkout", http.StatusSeeOther)
	}
}
