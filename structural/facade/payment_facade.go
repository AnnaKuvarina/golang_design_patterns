package main

import (
	"fmt"
	"time"
)

// Subsystem 1: Payment
type Payment struct{}

func (p *Payment) ProcessPayment(amount float64) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Processing payment of $%.2f\n", amount)
}

// Subsystem 2: Inventory
type Inventory struct{}

func (i *Inventory) CheckStock(itemID string) bool {
	fmt.Printf("Checking stock for item %s\n", itemID)
	// Simulate that the item is in stock
	return true
}

func (i *Inventory) ReserveItem(itemID string) {
	fmt.Printf("Reserving item %s\n", itemID)

}

// Subsystem 3: Shipping
type Shipping struct{}

func (s *Shipping) ArrangeShipping(itemID string) {
	fmt.Printf("Arranging shipping for item %s\n", itemID)
}

// Facade: Order Processor
type OrderProcessorFacade struct {
	payment   *Payment
	inventory *Inventory
	shipping  *Shipping
}

func NewOrderProcessorFacade(payment *Payment, inventory *Inventory, shipping *Shipping) *OrderProcessorFacade {
	return &OrderProcessorFacade{
		payment:   payment,
		inventory: inventory,
		shipping:  shipping,
	}
}

func (o *OrderProcessorFacade) ProcessOrder(itemID string, amount float64) {
	fmt.Println("Starting order processing...")
	if o.inventory.CheckStock(itemID) {
		o.payment.ProcessPayment(amount)
		o.inventory.ReserveItem(itemID)
		o.shipping.ArrangeShipping(itemID)
	} else {
		fmt.Println("Item is out of stock!")
	}
}
