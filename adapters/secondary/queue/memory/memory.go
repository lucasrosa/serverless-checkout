package queueadaptermemory

import (
	"fmt"

	"github.com/lucasrosa/serverless-checkout/businesslogic/cart"
)

type checkoutRepository struct{}

func NewMemoryCheckoutRepository() cart.CheckoutSecondaryPort {
	return &checkoutRepository{}
}

func (r *checkoutRepository) SendOrderForProcessing(order *cart.Order) error {
	fmt.Println(order)
	return nil
}
