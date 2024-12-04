package request_model

type Money struct {
	Coins1   int64 `json:"coins1"`
	Coins5   int64 `json:"coins5"`
	Coins10  int64 `json:"coins10"`
	Bank20   int64 `json:"bank20"`
	Bank50   int64 `json:"bank50"`
	Bank100  int64 `json:"bank100"`
	Bank500  int64 `json:"bank500"`
	Bank1000 int64 `json:"bank1000"`
}

type Checkout struct {
	ProductId string `json:"productId"`
	Total     Money  `json:"total"`
}