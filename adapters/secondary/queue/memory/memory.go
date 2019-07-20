package queueadaptermemory

import (
	"fmt"

	"github.com/lucasrosa/serverless-checkout/businesslogic/cart"
)

type checkoutRepository struct{}

// NewMemoryCheckoutRepository instantiates the repository for this adapter
func NewMemoryCheckoutRepository() cart.CheckoutSecondaryPort {
	return &checkoutRepository{}
}

func (r *checkoutRepository) SendOrderForProcessing(order *cart.Order) error {
	fmt.Println(order)
	return nil
}
