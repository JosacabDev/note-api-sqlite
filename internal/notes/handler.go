package notes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi/v5"
)

type HandlerNote struct {
	Repo Repository
}

func NewHandlerNote(repo Repository) *HandlerNote {
	return &HandlerNote{
		Repo: repo,
	}
}

func (h *HandlerNote) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	// Handler for listing all notes
	notes, err := h.Repo.GetAllNotes()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve notes. Error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(notes)
}

func (h *HandlerNote) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}
	noteID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID format", http.StatusBadRequest)
		return
	}
	note, err := h.Repo.GetNoteByID(noteID)
	if err != nil {
		if err.Error() == "note not found" {
			http.Error(w, "Note not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Failed to retrieve note. Error: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(note)
}

func (h *HandlerNote) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.Repo.CreateNote(note)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create note. Error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (h *HandlerNote) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}
	noteID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID format", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	note.ID = noteID

	err = h.Repo.UpdateNote(note)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update note. Error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HandlerNote) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}
	noteID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID format", http.StatusBadRequest)
		return
	}

	err = h.Repo.DeleteNote(noteID)
	if err != nil {

		http.Error(w, fmt.Sprintf("Failed to delete note. Error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
