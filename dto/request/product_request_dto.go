package request

// ProductRequest is used to bind data from client requests for creating a product
type ProductRequest struct {
	ID    uint   `json:"ID"`
	Name  string `json:"name"`
	Price string `json:"price"`
}
