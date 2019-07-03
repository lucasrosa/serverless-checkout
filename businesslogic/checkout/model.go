package checkout

// Order represents a buying order from a customer
type Order struct {
	Email        string  `json:"email"`
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	Description  string  `json:"description"`
	PaymentToken string  `json:"paymenttoken"`
}
