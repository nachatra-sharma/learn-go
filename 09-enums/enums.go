package main

import "fmt"

type OrderStatus string

const (
	RECEIVED = "received"
	CONFIRMED = "confirmed"
	PREPARED = "prepared"
	DELIVERED = "delivered"
)

func order(status OrderStatus) {
	fmt.Println("Order has been successfully: ", status)
}

func main() {
	order(PREPARED)
}
