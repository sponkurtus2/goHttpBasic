package monster

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("Received request to make a monster ^^")
	_, err := w.Write([]byte("Monster created"))
	if err != nil {
		log.Fatal("Error at creating monster :o")
	}
}

func (h *Handler) FindByMonsterID(w http.ResponseWriter, r *http.Request) {
	log.Println("Finding by id...")

	monster, exists := loadMonstersById()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(monster)
	if err != nil {
		log.Fatal("Error when encoding monster")
	}
}

/*
func (h *Handler) FindByMonsterName(w http.ResponseWriter, r *http.Request) {
	log.Println("Finding by name...")

	monster, exists := loadMonstersByName()[r.PathValue("name")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(monster)
	if err != nil {
		log.Fatal("Error when encoding monster")
	}
}
*/
