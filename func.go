package main

import (
	"fmt"
)

type Tho struct {
	id   int
	name int
	age  int
}

func (t Tho) anCo() {
	fmt.Println("Tho dang choi co")
}

func ham() {

	tho1 := Tho{}
	tho1.anCo()
}
