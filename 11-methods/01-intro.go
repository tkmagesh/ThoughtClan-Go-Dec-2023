package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	var pen Product
	// fmt.Println(pen)
	// fmt.Printf("%#v\n", pen)
	pen.Id = 100
	pen.Name = "Pen"
	pen.Cost = 10
	// fmt.Printf("%+v\n", pen)
	// fmt.Println(Format(pen))
	pen.WhoAmI()
	fmt.Println(pen.Format())

	var pencil Product = Product{
		Id:   101,
		Name: "Pencil",
	}
	// fmt.Printf("%+v\n", pencil)
	// fmt.Println(Format(pencil))
	fmt.Println(pencil.Format())

	var marker = Product{102, "Marker", 50} // Not advisable
	// fmt.Println(Format(marker))
	fmt.Println(marker.Format())

	fmt.Println("After applying 10% discount on marker")
	// ApplyDiscount(&marker, 10)
	marker.ApplyDiscount(10)
	// fmt.Println(Format(marker))
	fmt.Println(marker.Format())

	// pointer
	var markerPtr *Product = &marker
	fmt.Println((*markerPtr).Id)
	fmt.Println(markerPtr.Id)

	stapler := NewProduct(104, "Stapler", 75)

	// Have to dereference when passed as an argument to the function
	// fmt.Println(Format(*stapler))

	// No need to dereference when the function is invoked as a method
	fmt.Println(stapler.Format())

}

// Factory function for the product
// Encapsulates the complexity of constructing a product object
func NewProduct(id int, name string, cost float32) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

func (Product) WhoAmI() {
	fmt.Println("I am a product!")
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}

func (product *Product) ApplyDiscount(discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
