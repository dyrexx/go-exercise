package main

import "fmt"

func pointer() {
	x := 8

	// y := &x

	// fmt.Println(*y)

	// *y = 4

	// fmt.Println(*y)

	change(&x)

	fmt.Println(x)
}

func change(a *int) {
	*a = 6
}
