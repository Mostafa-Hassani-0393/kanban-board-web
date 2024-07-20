package card

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"kanban-board-app/internal/models"
	"net/http"
)

func CreateCard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var card models.Card
		if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Create(&card).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(card)
	}
}

func GetCards(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cards []models.Card
		if err := db.Find(&cards).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(cards)
	}
}

func UpdateCard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var card models.Card
		vars := mux.Vars(r)
		id := vars["id"]
		if err := db.First(&card, id).Error; err != nil {
			http.Error(w, "Card not found", http.StatusNotFound)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Save(&card).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(card)
	}
}

func DeleteCard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var card models.Card
		vars := mux.Vars(r)
		id := vars["id"]
		if err := db.Delete(&card, id).Error; err != nil {
			http.Error(w, "Card not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
