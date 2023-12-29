package main

import "fmt"

func main() {
	var pen struct {
		Id   int
		Name string
		Cost float32
	}
	// fmt.Println(pen)
	// fmt.Printf("%#v\n", pen)
	pen.Id = 100
	pen.Name = "Pen"
	pen.Cost = 10
	// fmt.Printf("%+v\n", pen)
	fmt.Println(Format(pen))

	var pencil struct {
		Id   int
		Name string
		Cost float32
	} = struct {
		Id   int
		Name string
		Cost float32
	}{
		Id:   101,
		Name: "Pencil",
		Cost: 5,
	}
	// fmt.Printf("%+v\n", pencil)
	fmt.Println(Format(pencil))
}

func Format(product struct {
	Id   int
	Name string
	Cost float32
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}
