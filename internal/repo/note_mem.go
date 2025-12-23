package repo

import (
	"sort"
	"sync"

	"example.com/notes-api/internal/core"
)

type NoteRepoMem struct {
	mu    sync.Mutex
	notes map[int64]core.Note
	next  int64
}

func NewNoteRepoMem() *NoteRepoMem {
	return &NoteRepoMem{notes: make(map[int64]core.Note)}
}

func (r *NoteRepoMem) Create(n core.Note) (core.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.next++
	n.ID = r.next
	r.notes[n.ID] = n
	return n, nil
}

func (r *NoteRepoMem) List() ([]core.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	out := make([]core.Note, 0, len(r.notes))
	for _, n := range r.notes {
		out = append(out, n)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

func (r *NoteRepoMem) Get(id int64) (core.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	n, ok := r.notes[id]
	if !ok {
		return core.Note{}, core.ErrNotFound
	}
	return n, nil
}

func (r *NoteRepoMem) Update(n core.Note) (core.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.notes[n.ID]; !ok {
		return core.Note{}, core.ErrNotFound
	}
	r.notes[n.ID] = n
	return n, nil
}

func (r *NoteRepoMem) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.notes[id]; !ok {
		return core.ErrNotFound
	}
	delete(r.notes, id)
	return nil
}
