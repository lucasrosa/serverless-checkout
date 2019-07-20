package cart

// CheckoutPrimaryPort is the entrypoint for the checkout Package
type CheckoutPrimaryPort interface {
	PlaceOrder(order *Order) error
}

// CheckoutSecondaryPort is the way the business rules communicate to the external world
type CheckoutSecondaryPort interface {
	SendOrderForProcessing(order *Order) error
}

// ProcessPrimaryPort is the entrypoint for the checkout Package
type ProcessPrimaryPort interface {
	ProcessOrder(order *Order) error
}

// ProcessSecondaryPort is the way the business rules communicate to the external world
type ProcessSecondaryPort interface {
	Save(order *Order) error
}
