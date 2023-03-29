package api

import (
	"github.com/gorilla/mux"
	"github.com/timur-danilchenko/go-nodes/api/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/notes", handlers.CreateNoteHandler).Methods("POST")
	router.HandleFunc("/notes", handlers.GetAllNotesHandler).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.GetNoteByIDHandler).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.UpdateNoteHandler).Methods("PUT")
	router.HandleFunc("/notes/{id}", handlers.DeleteNoteHandler).Methods("DELETE")
}
