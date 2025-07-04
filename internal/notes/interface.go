package notes

type Repository interface {
	GetAllNotes() ([]Note, error)
	GetNoteByID(id int64) (*Note, error)
	CreateNote(note Note) (*Note, error)
	UpdateNote(note Note) (*Note, error)
	DeleteNote(id int64) error
}
