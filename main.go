package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("name: %s, age: %d, money: %d, grades: %f, happiness: %f", u.Name, u.Age, u.Money, u.Avg_grades, u.Happiness)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Dancing"}}
	// fmt.Fprintf(page, `<h1> Main text </h1> <b> main text </b>`) // куда выводить и что выводить
	tmpl, _ := template.ParseFiles("templates/home_page.html") // откуда загружаем шаблон
	tmpl.Execute(w, bob)                                       // страничка и что передаем внутрь шаблона
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
