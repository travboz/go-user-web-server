package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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
	mux.HandleFunc("GET /users", app.getAllUsers)
	mux.HandleFunc("GET /users/{id}", app.getUserById)
	mux.HandleFunc("DELETE /users/{id}", app.deleteUserById)

	addr := ":4545"

	log.Printf("Server running on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
