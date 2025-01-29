package main

import (
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Name string `json:"name"`
}

type application struct {
	cache Storage
}

func main() {
	mux := http.NewServeMux()
	storage := NewSafeCache()

	app := application{
		cache: storage,
	}

	mux.HandleFunc("GET /", app.helloHandler)
	mux.HandleFunc("POST /users", app.createUser)
	mux.HandleFunc("GET /users/{id}", app.getUserById)
	mux.HandleFunc("DELETE /users/{id}", app.deleteUserById)

	addr := ":4545"

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
