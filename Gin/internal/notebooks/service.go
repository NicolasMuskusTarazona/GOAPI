package notebooks

import "fmt"

type NoteBookService struct {
	repo NoteBookRepository
}

func NewNoteBookService(r NoteBookRepository) *NoteBookService {
	return &NoteBookService{repo: r}
}
// Obtener todos
func (s*NoteBookService)GetAll()([]NoteBook, error){
	return  s.repo.FindAll()
}
// Obtener por ID
func (s*NoteBookService)GetByID(id int)(*NoteBook, error){
	return s.repo.FindByID(id)
}
// Crear
func (s *NoteBookService) CreateNoteBook(b NoteBook) error {
	if b.Stock < 0 || b.Price < 0 {
		return fmt.Errorf("stock cannot be negative")
	}
	return s.repo.Create(b)
}
// Actualizar
func (s *NoteBookService)UpdateNoteBook(b NoteBook)error  {
	return s.repo.Update(b)
}
// Eliminar
func (s*NoteBookService)DeleteNoteBook(id int)error  {
	return s.repo.Delete(id)
}