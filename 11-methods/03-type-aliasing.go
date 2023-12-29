package main

import "fmt"

type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	str := MyString("Proident ex officia eiusmod officia occaecat aliquip qui. Cupidatat et duis ut mollit ut reprehenderit amet. Sit irure dolore aliqua commodo aliquip aliqua do ut amet magna labore veniam sit nisi. Ut culpa nisi culpa eiusmod adipisicing qui quis occaecat esse adipisicing nulla. Tempor adipisicing ipsum irure et eiusmod pariatur ex irure id ad laboris occaecat ullamco anim. Amet cupidatat nostrud veniam eiusmod consectetur eiusmod.")
	fmt.Println(str.Length())
}
