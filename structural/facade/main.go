package main

func main() {
	payment := &Payment{}
	inventory := &Inventory{}
	shipping := &Shipping{}
	facade := NewOrderProcessorFacade(payment, inventory, shipping)
	facade.ProcessOrder("12345", 299.99)
}
