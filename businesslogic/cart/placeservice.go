package cart

type port struct {
	repo CheckoutSecondaryPort
}

// NewCheckoutService receives a Secondary Port of domain and insantiates a Primary Port
func NewCheckoutService(repo CheckoutSecondaryPort) CheckoutPrimaryPort {
	return &port{
		repo,
	}
}

func (p *port) PlaceOrder(order *Order) error {
	err := p.repo.SendOrderForProcessing(order)
	return err
}
