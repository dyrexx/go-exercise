package main

import "fmt"

func main() {
	port := [65535]bool{}
	port[0] = true
	fmt.Println(port[0])
}
