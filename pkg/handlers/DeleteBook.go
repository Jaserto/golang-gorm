package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go/crud/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock Books
	for index, book := range mocks.Books {
		if book.Id == id {
			// Delete book and send response if the book Id matches dynamic Id
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
