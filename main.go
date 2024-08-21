package main

import (
	"fmt"
	"net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Go is super easy!") // куда выводить и что выводить
}

// http.ResponseWriter - параметр, позволяющий выводить что-то на страничку
// http.Request - параметр, отвечающий за запрос

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page) // переходит к методу при попадании на определенный url адрес
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil) // отслеживает порт
}

func main() {
	handleRequest()
}
