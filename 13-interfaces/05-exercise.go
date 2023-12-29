package main

import (
	"fmt"
	"sort"
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

/* Sorting */

func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

// sort by Name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

// sort by Cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

// sort by Units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

// sort by Category
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	default:
		sort.Sort(products)
	}
}

func (products Products) Sort2(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(products, func(i int, j int) bool {
			return products[i].Id < products[j].Id
		})
	case "Name":
		sort.Slice(products, func(i int, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i int, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i int, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i int, j int) bool {
			return products[i].Category < products[j].Category
		})
	default:
		sort.Sort(products)
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

	fmt.Println("Sort by Id")
	products.Sort2("Id")
	// sort.Sort(products)
	fmt.Println(products)

	fmt.Println("Sort by Name")
	// sort.Sort(ByName{products})
	products.Sort2("Name")
	fmt.Println(products)

	fmt.Println("Sort by Cost")
	// sort.Sort(ByCost{products})
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println("Sort by Units")
	// sort.Sort(ByUnits{products})
	products.Sort("Units")
	fmt.Println(products)

	fmt.Println("Sort by Category")
	// sort.Sort(ByCategory{products})
	products.Sort("Category")
	fmt.Println(products)
}
