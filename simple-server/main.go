package main

import (
	"fmt"
	"net/http"
)

type names struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {

	http.HandleFunc("/", simpleServer)
	http.HandleFunc("/names", getNames)
	http.ListenAndServe(":8080", nil)
}
func simpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}
func getNames(w http.ResponseWriter, r *http.Request) {
	var firstName string
	var lastName string
	fmt.Fprintf(w, "Enter your first name: %s", r.URL.Path[1:])
	fmt.Scan(&firstName)
	fmt.Fprintf(w, "Enter your last name: %s", r.URL.Path[1:])
	fmt.Scan(&lastName)

}
