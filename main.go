package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Client struct {
	Name                  string
	Age                   uint16
	Balance               int16
	Happiness, Avg_grades float64
	Cars                  []string
}

func (c Client) printAllInfo() string {
	return fmt.Sprintf("Actually, Client name is %s, and he is %d age of old,"+
		"and as I know he have %d in balance", c.Name, c.Age, c.Balance)
}

func (c *Client) setNewName(newName string) {
	c.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	Azat := Client{"Azat", 24, 5000, 5.3, 5.4, []string{"Porsche 911", "G-wagon", "Ferrari"}}
	// Azat.setNewName("Naruto")
	// fmt.Fprintf(w, Azat.printAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, Azat)

}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
