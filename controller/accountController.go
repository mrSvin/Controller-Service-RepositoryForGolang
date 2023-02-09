package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"postgresql/service"
)

type Response struct {
	Message string `json:"message"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func AccountController() {
	http.HandleFunc("/userCreate", accountCreate)
	http.HandleFunc("/userRead", accountRead)
	http.HandleFunc("/userUpdate", accountUpdate)
	http.HandleFunc("/userDelete", accountDelete)

	http.ListenAndServe(":8080", nil)
}

func accountCreate(w http.ResponseWriter, r *http.Request) {
	data := Response{}
	if r.Method == "POST" {
		var user User
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		service.AccountCreate(user.Name, user.Password, user.Email)
		data = Response{
			Message: "Account create",
		}
		w.Header().Set("Content-Type", "application/json")

	} else {
		data = Response{
			Message: "need POST request",
		}
	}
	json.NewEncoder(w).Encode(data)
}

func accountRead(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")

		name, email := service.AccountRead(id)
		fmt.Println(name)
		fmt.Println(email)
		type UserRead struct {
			Name  string `json:"password"`
			Email string `json:"email"`
		}

		data := UserRead{
			Name:  name,
			Email: email,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		data := Response{
			Message: "need GET request",
		}
		json.NewEncoder(w).Encode(data)
	}

}

func accountUpdate(w http.ResponseWriter, r *http.Request) {
	data := Response{}
	if r.Method == "PUT" {
		var user User
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		id := r.URL.Query().Get("id")

		service.AccountUpdate(id, user.Name, user.Password, user.Email)

		data = Response{
			Message: "Account update",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		data = Response{
			Message: "need POST request",
		}
	}
}

func accountDelete(w http.ResponseWriter, r *http.Request) {
	data := Response{}
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		service.AccountDelete(id)

		data = Response{
			Message: "Account delete",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	} else {
		data = Response{
			Message: "need DELETE request",
		}
	}
}
