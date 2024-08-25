package main

import "fmt"

func main() {
	// Create different payment strategies
	creditCard := &CreditCardPayment{
		Name:        "John Doe",
		CardNumber:  "1234-5678-9101-1121",
		CVV:         "123",
		ExpiryMonth: 12,
		ExpiryYear:  2025,
	}
	paypal := &PayPalPayment{
		Email: "john.doe@example.com",
	}
	bitcoin := &BitcoinPayment{
		WalletAddress: "1A2b3C4d5E6f7G8h9I0jK",
	}

	processor := PaymentProcessor{}
	// Process payment using Credit Card
	processor.SetStrategy(creditCard)
	fmt.Println(processor.strategy.Pay(2.14))

	// Process payment using PayPal
	processor.SetStrategy(paypal)
	fmt.Println(processor.strategy.Pay(5.25))

	// Process payment using Bitcoin
	processor.SetStrategy(bitcoin)
	fmt.Println(processor.strategy.Pay(3.9))
}
