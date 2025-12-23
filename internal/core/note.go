package core

import "time"

type Note struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type NoteRepository interface {
	Create(n Note) (Note, error)
	List() ([]Note, error)
	Get(id int64) (Note, error)
	Update(n Note) (Note, error)
	Delete(id int64) error
}

var (
	ErrNotFound     = errorString("not found")
	ErrInvalidInput = errorString("invalid input")
)

type errorString string

func (e errorString) Error() string { return string(e) }

type NoteCreate struct {
	Title   string `json:"title" example:"Новая заметка"`
	Content string `json:"content" example:"Текст заметки"`
}

type NoteUpdate struct {
	Title   *string `json:"title,omitempty" example:"Обновлено"`
	Content *string `json:"content,omitempty" example:"Новый текст"`
}
