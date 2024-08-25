package main

// Context: Payment Processor
type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}
