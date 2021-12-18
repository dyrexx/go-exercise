package main

import "fmt"

func swap2() {
	a, b := 4, 5
	fmt.Println("Before: ", a, b)
	// Su dung bien tam

	// tmp := b
	// b = a
	// a = tmp
	// fmt.Println("After: ", a, b)

	// Khong su dung bien tam
	// b = b + a
	// a = b - a
	// b = b - a

	// fmt.Println("After: ", a, b)

	// Su dung pointer
	fmt.Println(swap(&a, &b))
}

func swap(a, b *int) (int, int) {
	tmp := *a
	*a = *b
	b = &tmp
	return *a, *b
}
