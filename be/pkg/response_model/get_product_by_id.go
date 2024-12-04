package response_model

type Product struct {
	ProductId string `json:"productId"`
	Name      string `json:"name"`
	ImageUrl  string `json:"imageUrl"`
	Stock     int64  `json:"stock"`
	Price     int64  `json:"price"`
}
