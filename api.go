package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var data []string = []string{}

// TestHandlerFunc is just a testing and learning handler func.
func TestHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string
	}{
		ID: "123",
	})
}

// AddItemsHandlerFunc adds items to a slice of strings
func AddItemsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	/*
	   Access the dynamic variable
	   Vars returns a map with keys(Var identifier) and value as strings
	   Here we have access to item variable value
	*/
	routeVar := mux.Vars(r)["item"]
	data = append(data, routeVar)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/test", TestHandlerFunc)
	r.HandleFunc("/add/{item}", AddItemsHandlerFunc).Methods("POST")

	http.ListenAndServe(":5000", r)
}
