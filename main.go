package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"user_id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

var users = []User{
	{
		ID:       1,
		Name:     "Bastian",
		Email:    "bastian@gmail.com",
		Password: "123456",
	},
	{
		ID:       2,
		Name:     "Lontsi",
		Email:    "lontsi@gmail.com",
		Password: "123lontsi456",
	},
	{
		ID:       3,
		Name:     "Kemka",
		Email:    "kemka@gmail.com",
		Password: "12kemka3456",
	},
}

func HalloFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo world")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	//Retourn responce
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.Write(resp)

}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HalloFunction)
	r.HandleFunc("/users", GetUsers).Methods("GET")

	fmt.Print("Server listen on port 9000")
	err := http.ListenAndServe(":9000", r)
	if err != nil {
		log.Fatal(err)
	}
}
