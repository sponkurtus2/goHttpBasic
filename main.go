package main

import (
	"log"
	"net/http"

	"github.com/sponkurtus2/goHttpBasic/monster"
)

// Path parameters
func getID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := w.Write([]byte("ID: " + id + "\n"))
	if err != nil {
		log.Fatal("Error at id :o")
	}
}

// Methods
func getName(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Handle GET
		_, err := w.Write([]byte("GET METHOD"))
		if err != nil {
			log.Fatal("Error at id :o")
		}
	case http.MethodPost:
		// Handle POST
		_, err := w.Write([]byte("POST METHOD"))
		if err != nil {
			log.Fatal("Error at id :o")
		}
	case http.MethodPut:
		// Handle PUT
		_, err := w.Write([]byte("PUT METHOD"))
		if err != nil {
			log.Fatal("Error at id :o")
		}
	}
}

func main() {
	router := http.NewServeMux()

	handler := &monster.Handler{}

	router.HandleFunc("POST /monster", handler.Create)
	router.HandleFunc("GET /monster/{id}", handler.FindByMonsterID)
	// router.HandleFunc("GET /monster/{id}", handler.FindByMonsterName)

	router.HandleFunc("/item/{id}", getID)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started at 8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error at server :O")
	}
}
