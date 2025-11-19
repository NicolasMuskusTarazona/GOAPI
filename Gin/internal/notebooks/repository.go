package notebooks

import "fmt"

type NoteBookRepository interface {
	FindAll() ([]NoteBook, error)
	FindByID(id int) (*NoteBook, error)
	Create(notebook NoteBook) error
	Update(notebook NoteBook) error
	Delete(id int) error
}

type InMemoryNoteBookRepository struct {
	data []NoteBook
}

func NewInMemoryNoteBookRepository()*InMemoryNoteBookRepository  {
	return &InMemoryNoteBookRepository{
		data: []NoteBook{},
	}
}
// OBTENER TODOS
func (r *InMemoryNoteBookRepository) FindAll() ([]NoteBook,error)  {
	return r.data, nil
}
// CREAR
func (r *InMemoryNoteBookRepository) Create(notebook NoteBook) error {
	notebook.ID = len(r.data) + 1
	r.data = append(r.data, notebook)
	return nil
}
// ACTUALIZAR
func (r*InMemoryNoteBookRepository) Update(notebook NoteBook) error {
	for i,b :=range r.data{
		if b.ID == notebook.ID{
			r.data[i] = notebook
			return nil
		}
	}
	return fmt.Errorf("notebook not found")
}
// Obtener por ID
func (r *InMemoryNoteBookRepository) FindByID(id int) (*NoteBook, error) {
	for _, b := range r.data {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, fmt.Errorf("notebook not found")
}
// Eliminar
func (r *InMemoryNoteBookRepository) Delete(id int)error {
	for i, b := range r.data{
		if b.ID == id{
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("notebook not found")
}
