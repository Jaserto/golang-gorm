package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/go/crud/pkg/mocks"
	"github.com/go/crud/pkg/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	//Read to request de Body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)
	//Append to the books mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	//Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
