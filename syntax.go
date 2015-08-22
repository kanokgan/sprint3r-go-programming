package main

import "fmt"


// ตัวอย่างเรื่อง function
func add(a, b int) int {
	return a + b
}

var addFunc func(int, int) int = func(a, b int) int {
	return a + b
}


// ตัวอย่างสร้าง Framework เล็ก ๆ เอาเอง
var router map[string]func(data string) = make(map[string]func(string))

func SetRoute(route string, f func(data string)) {
	router[route] = f
}

func RunRoute(route string, msg string) {
	if f, ok := router[route]; ok {
		f(msg)
	}
}


func main() {
	fmt.Println("10 + 20 =", addFunc(10, 20))
	fmt.Println("20 + 30 =", add(20, 30))

	SetRoute("/hello", func(msg string) {
		fmt.Println("Hello", msg)
	})

	RunRoute("/hello", "Kanokgan")
}