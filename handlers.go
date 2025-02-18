package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// helloHandler - works as a health checker GET.
// Returns some plaintext.
func (a *application) helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello there!\nwe're running on %s.", a.config.env)
}

// createUser - handles POST requests to create a new user.
// It expects a JSON object of the form `{ "name": "bob" }` and returns a JSON object of the shape.
func (a *application) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Name == "" {
		http.Error(w, ErrNameRequired.Error(), http.StatusBadRequest)
		return
	}

	user = a.cache.Insert(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// getAllUsers - handles GET requests to retrieve all users.
// It returns a JSON array of users.
func (a *application) getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users := a.cache.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// getUserById - handles GET requests to fetch a user by their ID.
// It expects an integer ID in the URL path and returns a JSON object.
func (a *application) getUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.cache.Get(id)

	if err != nil {
		http.Error(w, ErrUserDoesNotExist.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// deleteUserById - handles DELETE requests to delete a user by their ID.
// It expects an integer ID in the URL path and returns nothing.
func (a *application) deleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.cache.Delete(id)
	if err != nil {
		http.Error(w, "user does not exist", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
