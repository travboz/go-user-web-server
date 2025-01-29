package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (a *application) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there!")
}

func (a *application) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is required", http.StatusInternalServerError)
		return
	}

	user = a.cache.Insert(user)

	fmt.Fprintf(w, "The user's name is: %s", user.Name)
	w.WriteHeader(http.StatusNoContent)
}

func (a *application) getUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.cache.Get(id)

	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
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
