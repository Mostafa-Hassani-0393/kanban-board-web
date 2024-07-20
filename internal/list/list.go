package list

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"kanban-board-app/internal/models"
	"net/http"
)

func CreateList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.List
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Create(&list).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(list)
	}
}

func GetLists(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var lists []models.List
		if err := db.Find(&lists).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(lists)
	}
}

func UpdateList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.List
		vars := mux.Vars(r)
		id := vars["id"]
		if err := db.First(&list, id).Error; err != nil {
			http.Error(w, "List not found", http.StatusNotFound)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(list)
	}
}

func DeleteList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list models.List
		vars := mux.Vars(r)
		id := vars["id"]
		if err := db.Delete(&list, id).Error; err != nil {
			http.Error(w, "List not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
