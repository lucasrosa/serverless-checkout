package cart

import "testing"

type processRepository struct{}

func (r *processRepository) Save(order *Order) error {
	return nil
}

func TestProcessOrder(t *testing.T) {
	order := Order{}
	repo := &processRepository{}
	service := NewProcessService(repo)
	err := service.ProcessOrder(&order)

	if err != nil {
		t.Errorf("ProcessOrder is not executing with correct behavior.")
	}
}
