package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/timur-danilchenko/go-nodes/database"
	"github.com/timur-danilchenko/go-nodes/models"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	sqlStatement := `INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id`

	id := 0
	err = db.QueryRow(sqlStatement, note.Title, note.Content).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	note.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// GetAllNotesHandler retrieves all notes from the database
func GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// query the database
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// create a slice of notes to hold the results
	var notes []models.Note

	// iterate over the rows and scan the results into a note object
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		notes = append(notes, note)
	}

	// check for errors after iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the notes as a JSON array
	json.NewEncoder(w).Encode(notes)
}

func GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	note := &models.Note{}
	err = db.QueryRow("SELECT id, title, content FROM notes WHERE id = $1", noteID).Scan(&note.ID, &note.Title, &note.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve note", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(note)
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE notes SET title=$1, content=$2 WHERE id=$3", updatedNote.Title, updatedNote.Content, noteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the note to delete from the URL parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the note from the database
	_, err = db.Exec("DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Unable to delete note", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusNoContent)
}
