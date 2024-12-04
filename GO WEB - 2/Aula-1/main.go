package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)


type Greetings struct {
	FirstName	string `json: "firstName"`
	LastName 	string `json: "firstName"`
}

func main() {
	router := chi.NewRouter()
//Exercicio 1 
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Write([]byte(`pong`))
	})

//Exercicio 2
	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var requestedBody Greetings

		if err := json.NewDecoder(r.Body).Decode(&requestedBody); err != nil {
			panic(err)
		}
		
		responseBody := fmt.Sprintf("Hello %s %s", requestedBody.FirstName, requestedBody.LastName)

		w.WriteHeader(200)

		w.Write([]byte(responseBody))
	})


	

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}