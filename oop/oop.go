package main

import "fmt"

type Animal interface {
	eat()
}

type Cat struct {
	Id   int
	Name string
	Age  int
	Leg  int
	Tail bool
}

type Tho struct {
	Id   int
	Name string
	Age  int
	Leg  int
	Tail bool
}

func (c Cat) eat() {
	fmt.Println(c.Name, "Dang an kem")

}

func (t Tho) eat() {
	fmt.Println(t.Name, "Dang an co")

}

func main() {
	meoden := Cat{
		Id:   1,
		Name: "DuongHeo",
		Age:  11,
		Leg:  4,
		Tail: true,
	}

	thoden := Tho{1, "DuongHi", 12, 4, true}
	meoden.eat()
	thoden.eat()
	fmt.Println(meoden)
	fmt.Println(thoden)
}
