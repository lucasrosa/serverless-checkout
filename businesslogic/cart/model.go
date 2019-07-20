package cart

// Order represents a buying order from a customer
type Order struct {
	ID           string  `json:"id"`
	Email        string  `json:"email"`
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	ProductID    int     `json:"productid"`
	PaymentToken string  `json:"paymenttoken"`
}
