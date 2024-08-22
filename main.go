package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("name: %s, age: %d, money: %d, grades: %f, happiness: %f", u.name, u.age, u.money, u.avg_grades, u.happiness)
}

func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	fmt.Fprintf(page, bob.getAllInfo()) // куда выводить и что выводить
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
