package main

import "fmt"

func add(a, b int) int {
	return a + b
}

var addFunc func(int, int) int = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("10 + 20 =", addFunc(10, 20))
	fmt.Println("20 + 30 =", add(20, 30))
}