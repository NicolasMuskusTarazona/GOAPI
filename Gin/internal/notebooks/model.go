package notebooks

type NoteBook struct {
	ID 			int 	`json:"id,omitempty"`
	Price 		float64 `json:"price"`
	Stock 		int 	`json:"stock"`
}