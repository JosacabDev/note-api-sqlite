package notes

import (
	"database/sql"
	"errors"
)

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetAllNotes() ([]Note, error) {
	var notes []Note
	query := `SELECT id, title, content FROM notes`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) GetNoteByID(id int64) (*Note, error) {
	var note Note
	query := `SELECT id, title, content FROM notes WHERE id = ?`
	row := r.DB.QueryRow(query, id)
	if err := row.Scan(&note.ID, &note.Title, &note.Content); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("note not found")
		}
		return nil, err
	}
	return &note, nil
}

func (r *repository) CreateNote(note Note) (*Note, error) {
	query := `INSERT INTO notes (title, content) VALUES (?, ?)`
	result, err := r.DB.Exec(query, note.Title, note.Content)
	if err != nil {
		return nil, err
	}

	note.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdNote, err := r.GetNoteByID(note.ID)
	if err != nil {
		return nil, err
	}

	return createdNote, nil
}

func (r *repository) UpdateNote(note Note) (*Note, error) {
	query := `UPDATE notes SET title = ?, content = ? WHERE id = ?`
	_, err := r.DB.Exec(query, note.Title, note.Content, note.ID)
	if err != nil {
		return nil, err
	}

	updatedNote, err := r.GetNoteByID(note.ID)
	if err != nil {
		return nil, err
	}
	return updatedNote, nil
}

func (r *repository) DeleteNote(id int64) error {
	query := `DELETE FROM notes WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
