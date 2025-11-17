package books

import "fmt"

type BookService struct {
	repo BookRepository
}

func NewBookService(r BookRepository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) CreateBook(b Book) error {
	if b.Stock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}
	return s.repo.Create(b)
}

func (s *BookService)UpdateBook(b Book)error  {
	return s.repo.Update(b)
}