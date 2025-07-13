package notes

import (
	"encoding/json"
	"fmt"
	customErros "github/JosacabDev/api-sqlite/pkg/errors"
	"github/JosacabDev/api-sqlite/pkg/libjson"
	"github/JosacabDev/api-sqlite/pkg/logger"
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
	notes, err := h.Repo.GetAllNotes()
	if err != nil {
		logger.Error.Println("Failed to retrieve notes: ", err)
		libjson.EncodeCustomError(w, customErros.InternalError(fmt.Sprintf("Failed to retrieve notes. Error: %s", err.Error())))
		return
	}
	libjson.EncodeOk(w, notes)
}

func (h *HandlerNote) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		libjson.EncodeCustomError(w, customErros.BadRequestError("Note ID is required"))
		return
	}

	noteID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		libjson.EncodeCustomError(w, customErros.BadRequestError("Invalid note ID format"))
		return
	}

	note, err := h.Repo.GetNoteByID(noteID)
	if err != nil {
		logger.Error.Println("Failed to retrieve note: ", err)
		libjson.EncodeCustomError(w, err)
		return
	}

	libjson.EncodeOk(w, note)
}

func (h *HandlerNote) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		logger.Error.Println("Failed to decode request body:", err)
		libjson.EncodeCustomError(w, customErros.BadRequestError("Invalid request payload"))
		return
	}

	createdNote, err := h.Repo.CreateNote(note)
	if err != nil {
		logger.Error.Println("Failed to create note:", err)
		libjson.EncodeCustomError(w, customErros.InternalError(fmt.Sprintf("Failed to create note. Error: %s", err.Error())))
		return
	}

	libjson.EncodeCreated(w, createdNote)
}

func (h *HandlerNote) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	id := chi.URLParam(r, "id")
	if id == "" {
		logger.Error.Println("Note ID is required")
		libjson.EncodeCustomError(w, customErros.BadRequestError("Note ID is required"))
		return
	}
	noteID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logger.Error.Println("Invalid note ID format:", err)
		libjson.EncodeCustomError(w, customErros.BadRequestError("Invalid note ID format"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		logger.Error.Println("Failed to decode request body:", err)
		libjson.EncodeCustomError(w, customErros.BadRequestError("Invalid request payload"))
		return
	}
	note.ID = noteID

	updatedNote, err := h.Repo.UpdateNote(note)
	if err != nil {
		logger.Error.Println("Failed to update note:", err)
		libjson.EncodeCustomError(w, err)
		return
	}

	libjson.EncodeOk(w, updatedNote)
}

func (h *HandlerNote) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		logger.Error.Println("Note ID is required")
		libjson.EncodeCreated(w, customErros.BadRequestError("Note ID is required"))
		return
	}
	noteID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logger.Error.Println("Invalid note ID format:", err)
		libjson.EncodeCustomError(w, customErros.BadRequestError("Invalid note ID format"))
		return
	}

	err = h.Repo.DeleteNote(noteID)
	if err != nil {
		logger.Error.Println("Failed to delete note:", err)
		libjson.EncodeCustomError(w, customErros.InternalError(fmt.Sprintf("Failed to delete note. Error: %s", err.Error())))
		return
	}

	libjson.EncodeNoContent(w)
}
