package checkout

// PrimaryPort is the entrypoint for the checkout Package
type PrimaryPort interface {
	PlaceOrder(order *Order) error
}

// SecondaryPort is the
type SecondaryPort interface {
	SendOrderForProcessing(order *Order) error
}

// ProcessPrimaryPort is the entrypoint for the checkout Package
type ProcessPrimaryPort interface {
	ProcessOrder(order *Order) error
}

// ProcessSecondaryPort
type ProcessSecondaryPort interface {
	Save(order *Order) error
}
