package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"topup-game/models"
)

func RiwayatInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Mengambil nilai InvoiceNumber dari formulir HTTP
		InvoiceNumberStr := r.FormValue("invoiceNumber")
		InvoiceNumber, err := strconv.Atoi(InvoiceNumberStr)
		if err != nil {
			http.Error(w, "Invalid InvoiceNumber", http.StatusBadRequest)
			return
		}

		// Memanggil fungsi GetInvoiceByNumber dengan nilai InvoiceNumber yang sudah diubah
		result, err := models.GetInvoiceByNumber(InvoiceNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Menyiapkan dan mengeksekusi template HTML
		tmpl, err := template.ParseFiles("views/navbar/riwayat-invoice.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct{ Result *models.Invoice }{Result: result})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
