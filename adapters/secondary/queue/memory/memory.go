package queueadaptermemory

import (
	"fmt"

	"github.com/lucasrosa/serverless-checkout/businesslogic/checkout"
)

type checkoutRepository struct{}

func NewMemoryCheckoutRepository() checkout.SecondaryPort {
	return &checkoutRepository{}
}

func (r *checkoutRepository) SendOrderForProcessing(order *checkout.Order) error {
	fmt.Println(order)
	return nil
}
