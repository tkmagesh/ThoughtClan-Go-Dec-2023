package main

import (
	"fmt"
	"strings"
)

/*
Write an api for sorting the products by any attribute
IMPORTANT: MUST use sort.Sort()
*/

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

// fmt.Stringer interface implementation
func (products Products) String() string {
	var sb strings.Builder
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p))
	}
	return sb.String()
}

func (products Products) IndexOf(p Product) int {
	for idx, prod := range products {
		if prod == p {
			return idx
		}
	}
	return -1
}

/*
func (products Products) FilterCostlyProducts() Products {
	var result Products
	for _, p := range products {
		if p.Cost > 50 {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	var result Products
	for _, p := range products {
		if p.Category == "Stationary" {
			result = append(result, p)
		}
	}
	return result
}
*/

type ProductPredicate func(Product) bool

func (products Products) Filter(predicate ProductPredicate) Products {
	var result Products
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func negate(predicate ProductPredicate) ProductPredicate {
	return func(p Product) bool {
		return !predicate(p)
	}
}

func main() {

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products)

	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println("Index of stove :", products.IndexOf(stove))

	/*
		fmt.Println("Filter costly products")
		costlyProducts := products.FilterCostlyProducts()
		fmt.Println(costlyProducts)

		fmt.Println("Filter stationary products")
		stationaryProducts := products.FilterStationaryProducts()
		fmt.Println(stationaryProducts)
	*/

	fmt.Println("Filter costly products")
	var costlyProductPredicate = func(p Product) bool {
		return p.Cost > 50
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("Filter affordable products")
	/*
		var affordableProductPredicate = func(p Product) bool {
			return !costlyProductPredicate(p)
		}
	*/

	var affordableProductPredicate = negate(costlyProductPredicate)
	affordableProducts := products.Filter(affordableProductPredicate)
	fmt.Println(affordableProducts)

	fmt.Println("Filter stationary products")
	var stationaryProductPredicate = func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts)
}
