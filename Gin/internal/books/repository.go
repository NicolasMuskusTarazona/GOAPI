package books

import "fmt"

type BookRepository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (*Book, error)
	Create(book Book) error
	Update(book Book) error
	Delete(id int) error
}

type InMemoryBookRepository struct {
	data []Book
}

func NewInMemoryBookRepository()*InMemoryBookRepository  {
	return &InMemoryBookRepository{
		data: []Book{},
	}
}

func (r *InMemoryBookRepository) FindAll() ([]Book,error)  {
	return r.data, nil
}

func (r *InMemoryBookRepository) Create(book Book) error {
	book.ID = len(r.data) + 1
	r.data = append(r.data, book)
	return nil
}

func (r*InMemoryBookRepository) Update(book Book) error {
	for i,b :=range r.data{
		if b.ID == book.ID{
			r.data[i] = book
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

func (r *InMemoryBookRepository) FindByID(id int) (*Book, error) {
	for _, b := range r.data {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}

func (r *InMemoryBookRepository) Delete(id int)error {
	for i, b := range r.data{
		if b.ID == id{
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book not found")
}
