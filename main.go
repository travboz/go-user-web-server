package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	cache  Storage
	config config
}

type config struct {
	addr string
	env  string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mux := http.NewServeMux()
	storage := NewSafeCache()

	cfg := config{
		addr: GetString("SERVER_PORT", ":8081"),
		env:  GetString("ENV", "development"),
	}

	app := application{
		cache: storage,
	}

	mux.HandleFunc("GET /", app.helloHandler)
	mux.HandleFunc("POST /users", app.createUser)
	mux.HandleFunc("GET /users", app.getAllUsers)
	mux.HandleFunc("GET /users/{id}", app.getUserById)
	mux.HandleFunc("DELETE /users/{id}", app.deleteUserById)

	log.Printf("Server running on http://localhost%s", cfg.addr)
	if err := http.ListenAndServe(cfg.addr, mux); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
