package checkout

type port struct {
	repo SecondaryPort
}

// NewCheckoutService receives a Secondary Port of domain and insantiates a Primary Port
func NewCheckoutService(repo SecondaryPort) PrimaryPort {
	return &port{
		repo,
	}
}

func (p *port) PlaceOrder(order *Order) error {
	err := p.repo.SendOrderForProcessing(order)
	return err
}
