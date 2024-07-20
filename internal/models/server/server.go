package server

import (
	"github.com/gorilla/mux"
	"kanban-board-app/internal/list"

	"kanban-board-app/internal/auth"
	"kanban-board-app/internal/board"
	"kanban-board-app/internal/card"
	"kanban-board-app/pkg/database"

	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	//Authentication routes
	router.HandleFunc("/signup", auth.Signup).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	// Board routes
	router.HandleFunc("/boards", board.CreateBoard(database.DB)).Methods("POST")
	router.HandleFunc("/boards", board.GetBoards(database.DB)).Methods("GET")
	router.HandleFunc("/boards/{id}", board.UpdateBoard(database.DB)).Methods("PUT")
	router.HandleFunc("/boards/{id}", board.DeleteBoard(database.DB)).Methods("DELETE")

	// list routes
	router.HandleFunc("/lists", list.CreateList(database.DB)).Methods("POST")
	router.HandleFunc("/lists", list.GetLists(database.DB)).Methods("GET")
	router.HandleFunc("/lists/{id}", list.UpdateList(database.DB)).Methods("PUT")
	router.HandleFunc("/lists/{id}", list.DeleteList(database.DB)).Methods("DELETE")

	// card routes
	router.HandleFunc("/cards", card.CreateCard(database.DB)).Methods("POST")
	router.HandleFunc("/cards", card.GetCards(database.DB)).Methods("GET")
	router.HandleFunc("/cards/{id}", card.UpdateCard(database.DB)).Methods("PUT")
	router.HandleFunc("/cards/{id}", card.DeleteCard(database.DB)).Methods("DELETE")

	return router
}

func Run() {
	router := NewRouter()
	http.ListenAndServe(":8085", router)
}
