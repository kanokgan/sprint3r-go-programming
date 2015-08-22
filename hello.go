package main

import "fmt"

func main() {
	hello := "Hello World" // string
	fmt.Println(hello)

	var testInt int = 1 // int

	var testBool bool = true // bool

	var testfloat32 float32 = 1.0 // float32

	var testfloat64 float64 = 1.0 // float64

	var teststring string = "string" // string

	var testunit uint = 1 // uint

	var testint64 int64 = 1 // int64, long (in other language)

	var testbyte byte = 1 // byte

	var testCastStringtoByte []byte = []byte("Hello World") // cast string to byte array 

	fmt.Println(testInt)
	fmt.Println(testBool)
	fmt.Println(testfloat32)
	fmt.Println(testfloat64)
	fmt.Println(teststring)
	fmt.Println(testunit)
	fmt.Println(testint64)
	fmt.Println(testbyte)
	fmt.Println(testCastStringtoByte)
}