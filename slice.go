package main

import "fmt"

func slice() {
	// a := []int{1, 2, 5, 6, 0}

	// fmt.Printf("%T", a)

	//myNameIs := [5]string{"Hello", "My", "Name", "Is", "Khanh"}
	// Address of array muNameIS
	//omg := myNameIs[2:4]

	// Address of array slice pointer

	// omg = append(omg, "Truong", "Cute", "qua", "troi", "!", "!", "!")
	// fmt.Printf("After: %p", &omg)
	// fmt.Println(omg)
	// fmt.Println(cap(omg), len(omg))

	// Slice tham chiếu tới 1 array khác
	// Tạo 1 slice tham chiếu tới 1 array cụ thể thì slice có cap bằng 2 lần cap của array
	// Slice append thêm phần tử bằng cách tạo 1 array mới sau đó copy giá trị từ slice cũ sang + giá trị mới

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	c := append(a, b...)

	fmt.Println(c)

}
