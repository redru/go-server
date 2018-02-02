package goserver

import (
	"net/http"
)

// Default contructor
type Default struct {
}

// PersonsGet retrieve the list of Persons
func PersonsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
