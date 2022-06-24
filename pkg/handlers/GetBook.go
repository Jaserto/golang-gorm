package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go/crud/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id parameter

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//iterate over all the mock books

	for _, book := range mocks.Books {
		if book.Id == id {
			//if ids are equal send book as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
