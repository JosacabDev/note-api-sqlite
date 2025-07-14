package server

import (
	"database/sql"
	"github/JosacabDev/api-sqlite/internal/notes"
	"github/JosacabDev/api-sqlite/internal/tasks"
	"github/JosacabDev/api-sqlite/pkg/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port   string
	DB     *sql.DB
	Router *chi.Mux
}

func NewServer(port string, db *sql.DB) *Server {
	s := &Server{
		Port:   port,
		DB:     db,
		Router: chi.NewRouter(),
	}

	s.Router.Use(middleware.RequestLogger)
	s.setUpRoutes()
	return s
}

func (s *Server) setUpRoutes() {
	notesRepo := notes.NewRepository(s.DB)
	notesHandler := notes.NewHandlerNote(notesRepo)

	tasksRepo := tasks.NewRepository(s.DB)
	tasksHandler := tasks.NewHandlerTask(tasksRepo)

	// Notes
	s.Router.Route("/notes", func(r chi.Router) {
		r.Get("/", notesHandler.GetAllNotes)
		r.Get("/{id}", notesHandler.GetNoteByID)
		r.Post("/", notesHandler.CreateNote)
		r.Put("/{id}", notesHandler.UpdateNote)
		r.Delete("/{id}", notesHandler.DeleteNote)
	})

	// Tasks
	s.Router.Route("/tasks", func(r chi.Router) {
		r.Get("/", tasksHandler.GetAllTasks)
	})
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.Port, s.Router)
}
