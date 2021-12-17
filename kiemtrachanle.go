package main

import "fmt"

func main() {

	// Cách 1: chia lấy dư

	// for {
	// 	fmt.Print("Moi ban nhap so a:")
	// 	var a int
	// 	fmt.Scan(&a)

	// 	if a%2 == 0 {
	// 		fmt.Println(a, "la so chan")
	// 	} else {
	// 		fmt.Println(a, "so le")
	// 	}
	// }

	// Cách 2 dùng bitwise

	for {
		fmt.Print("Nhap a: ")
		var a int
		fmt.Scan(&a)
		if a&1 == 1 {
			fmt.Println("Số lẻ")
		} else {
			fmt.Println("Số chẵn")
		}
	}

}
