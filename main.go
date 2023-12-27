// main.go
package main

import (
	"log"
	"net/http"
	"topup-game/config"
	"topup-game/controller"
)

func main() {
	config.ConnectDB()

	//Assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Navigation
	http.HandleFunc("/", controller.Dashboard)
	http.HandleFunc("/contact", controller.Contact)
	http.HandleFunc("/riwayat", controller.Riwayat)
	http.HandleFunc("/riwayat-invoice", controller.RiwayatInvoiceHandler)

	//Game
	http.HandleFunc("/freefire", controller.Freefire)

	//Checkout
	http.HandleFunc("/checkout", controller.Checkout)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
