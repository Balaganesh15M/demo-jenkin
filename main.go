package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the User API! Visit /users to get the list.10-07-2025")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", getUsers)
	println("Starting server successfully...")
    println("webhook triggered successfully")

	log.Println("Server started at :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
