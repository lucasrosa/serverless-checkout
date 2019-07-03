package checkout

// PrimaryPort is the entrypoint for the checkout Package
type PrimaryPort interface {
	PlaceOrder(order *Order) error
}

// SecondaryPort is the
type SecondaryPort interface {
	SendOrderForProcessing(order *Order) error
}
