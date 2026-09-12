package main

import "fmt"

// payment method

type paymentMethod interface {
	pay(amount float32) bool
}

func makePayment(p paymentMethod, amount float32) {
	p.pay(amount)
}

// stripe
type stripe struct {
	apiKey string
	secretKey string
}

// constructor
func createStripe (apiKey string, secretKey string) *stripe {
	s := stripe{
		apiKey: apiKey,
		secretKey: secretKey,
	}
	return &s
}

// pay using stripe
func (s *stripe) pay(amount float32) bool {
	fmt.Println("Paid using stripe", amount)
	return true
}

// razorpay
type razorpay struct {
	apiKey string
	secretKey string
}

// constructor for razorpay
func createRazorpay(apiKey string, secretKey string) *razorpay {
	r := razorpay{
		apiKey: apiKey,
		secretKey: secretKey,
	}
	return &r
}

// pay using razorpay
func (r *razorpay) pay(amount float32) bool {
	fmt.Println("paid using razorpay !!", amount)
	return true
}

func main() {
	// pay := createRazorpay("123", "456")
	pay := createStripe("123", "456")
	makePayment(pay, 123)
}
