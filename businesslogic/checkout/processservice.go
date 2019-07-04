package checkout

type dbPort struct {
	repo ProcessSecondaryPort
}

// NewProcessService receives a Secondary Port of domain and insantiates a Primary Port
func NewProcessService(repo ProcessSecondaryPort) ProcessPrimaryPort {
	return &dbPort{
		repo,
	}
}

func (dp *dbPort) ProcessOrder(order *Order) error {
	err := dp.repo.Save(order)
	return err
}
