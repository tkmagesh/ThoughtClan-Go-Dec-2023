package main

import "fmt"

/*
Employee (id, name, org)
Organization (id, name, city)
*/
type Organization struct {
	Id   int
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	thoughtClan := &Organization{Id: 999, Name: "ThoughtClan", City: "Bengaluru"}
	e1 := Employee{Id: 100, Name: "Magesh", Org: thoughtClan}
	e2 := Employee{Id: 101, Name: "Suresh", Org: thoughtClan}

	fmt.Println(e1.Org.City, e2.Org.City)
	thoughtClan.City = "Mysuru"
	fmt.Println(e1.Org.City, e2.Org.City)
}
