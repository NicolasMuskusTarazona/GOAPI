package books 

type Book struct {
	ID 			int 	`json:"id,omitempty"`
	Title 		string 	`json:"title"`
	Author 		string  `json:"author"`
	Price 		float64 `json:"price"`
	Stock 		int 	`json:"stock"`
	Description string 	`json:"description"`
}