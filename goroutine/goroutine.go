package main

import (
	"fmt"
	"sync"
)

func f1() {
	fmt.Println("Day la f1")
	wait.Done()
}

func f2() {
	fmt.Println("Day la f2")
	wait.Done()
}

var wait = sync.WaitGroup{}

func main() {
	wait.Add(2)
	go f1()
	go f2()
	wait.Wait()
}
