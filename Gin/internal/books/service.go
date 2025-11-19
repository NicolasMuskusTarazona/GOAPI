package books

import "fmt"

type BookService struct {
	repo BookRepository
}

func NewBookService(r BookRepository) *BookService {
	return &BookService{repo: r}
}
func (s*BookService)GetAll()([]Book, error){
	return  s.repo.FindAll()
}

func (s*BookService)GetByID(id int)(*Book, error){
	return s.repo.FindByID(id)
}
func (s *BookService) CreateBook(b Book) error {
	if b.Stock < 0 || b.Price < 0 {
		return fmt.Errorf("stock cannot be negative")
	}
	return s.repo.Create(b)
}

func (s *BookService)UpdateBook(b Book)error  {
	return s.repo.Update(b)
}

func (s*BookService)DeleteBook(id int)error  {
	return s.repo.Delete(id)
}