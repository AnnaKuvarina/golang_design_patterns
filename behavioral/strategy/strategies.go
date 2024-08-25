package main

import "fmt"

// The strategy interface
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Concrete Strategy 1: Credit Card Payment
type CreditCardPayment struct {
	Name        string
	CardNumber  string
	CVV         string
	ExpiryMonth int
	ExpiryYear  int
}

func (cc *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card [%s]", amount, cc.CardNumber)
}

// Concrete Strategy 1: PayPal Payment
type PayPalPayment struct {
	Email string
}

func (pp *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal [%s]", amount, pp.Email)
}

type BitcoinPayment struct {
	WalletAddress string
}

func (b *BitcoinPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Bitcoin [%s]", amount, b.WalletAddress)
}
