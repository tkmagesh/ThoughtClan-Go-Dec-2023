package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	Dummy
	Product
	Expiry string
}

func main() {
	// var grapes PerishableProduct
	// var grapes PerishableProduct = PerishableProduct{Product{100, "Grapes", 60}, "2 Days"}
	var grapes PerishableProduct = PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Grapes",
			Cost: 60,
		},
		Expiry: "2 Days",
	}
	fmt.Println(grapes.Product.Id, grapes.Name, grapes.Cost, grapes.Expiry)
	// fmt.Println(grapes.Product.Id, grapes.Product.Name, grapes.Product.Cost, grapes.Expiry)

	// use the Format & ApplyDiscount functions for grapes (do not change the functions)
}

func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}

func ApplyDiscount(product *Product, discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
