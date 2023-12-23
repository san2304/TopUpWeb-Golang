package controller

import (
	"html/template"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Jika metodenya adalah GET, maka tampilkan halaman admin
		renderAdminPage(w)
	} else {
		// Jika metodenya bukan GET, tanggapi dengan metode tidak diizinkan
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func renderAdminPage(w http.ResponseWriter) {
	// Parsing template dari file
	temp, err := template.ParseFiles("views/admin/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Eksekusi template
	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
