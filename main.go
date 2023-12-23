// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"topup-game/config"
	"topup-game/controller"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Admin Panel - Port 4040")
}

func visitorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Selamat datang, Pengunjung! - Port 8080")
}

func main() {
	config.ConnectDB()

	// Handler untuk menyajikan aset (gambar, dll.)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// NAvigation
	http.HandleFunc("/", controller.Dashboard)
	http.HandleFunc("/contact", controller.Contact)
	http.HandleFunc("/invoice", controller.Invoice)
	http.HandleFunc("/check-invoice", controller.CheckInvoiceHandler)

	//game
	http.HandleFunc("/freefire", controller.Freefire)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
