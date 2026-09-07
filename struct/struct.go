package main

import "fmt"

type Product struct {
	name string
	price int
	company string
}

// this is a constructor function
func newProduct(name string, price int, company string) Product {
	p := Product{
		name: name,
		price: price,
		company: company,
	}
	return p
}

func (p *Product) displayProduct() {
	fmt.Println(*p)
}

func main() {
	p := newProduct("Iphone 16 pro max", 100000, "Apple inc.")
	p.name = "Macbook Pro"
	// p := Product{
	// 	name: "Iphone 16 pro max",
	// 	price: 1000000,
	// 	company: "Apple inc.",
	// }
	// fmt.Println(p)
	p.displayProduct()
}
