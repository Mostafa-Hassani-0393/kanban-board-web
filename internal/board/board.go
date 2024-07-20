package board

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"kanban-board-app/internal/models"
	"net/http"
)

func CreateBoard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var board models.Board
		if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Create(&board).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(board)
	}
}

func GetBoards(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var boards []models.Board
		if err := db.Find(&boards).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(boards)
	}
}

func UpdateBoard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var board models.Board
		vars := mux.Vars(r)
		id := vars["id"]
		if err := db.First(&board, id).Error; err != nil {
			http.Error(w, "Board not found", http.StatusNotFound)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Save(&board).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(board)
	}
}

func DeleteBoard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var board models.Board
		vars := mux.Vars(r)
		id := vars["id"]
		if err := db.Delete(&board, id).Error; err != nil {
			http.Error(w, "Board not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
