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
	fmt.Println(Format(pen))

	var pencil Product = Product{
		Id:   101,
		Name: "Pencil",
	}
	// fmt.Printf("%+v\n", pencil)
	fmt.Println(Format(pencil))

	var marker = Product{102, "Marker", 50} // Not advisable
	fmt.Println(Format(marker))

	fmt.Println("After applying 10% discount on marker")
	ApplyDiscount(&marker, 10)
	fmt.Println(Format(marker))

	// pointer
	var markerPtr *Product = &marker
	fmt.Println((*markerPtr).Id)
	fmt.Println(markerPtr.Id)

}

func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}

func ApplyDiscount(product *Product, discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}