package main

import "fmt"

func main() {
	HelloWorld := "Hello World" 		// string
	fmt.Println(HelloWorld)

	var testInt int = 1 				// int
	fmt.Println(testInt)

	var testBool bool = true 			// bool
	fmt.Println(testBool)

	var testfloat32 float32 = 1.0 		// float32
	fmt.Println(testfloat32)

	var testfloat64 float64 = 1.0 		// float64
	fmt.Println(testfloat64)

	var teststring string = "string" 	// string
	fmt.Println(teststring)

	var testunit uint = 1 				// uint
	fmt.Println(testunit)

	var testint64 int64 = 1 			// int64, long (in other language)
	fmt.Println(testint64)

	var testbyte byte = 1 				// byte
	fmt.Println(testbyte)

	var testCastStringtoByte []byte = []byte("Hello World") // cast string to byte array 
	fmt.Println(testCastStringtoByte)

	var testIntArray [6]int = [6]int{1, 2, 3, 4, 5} 		// array
	fmt.Println(testIntArray)

	testIntArray[1] = 9
	fmt.Println(testIntArray)

	var testSlice []int = []int{1, 2, 3, 4, 5, 6}

	testSlice = append(testSlice, 12, 3, 1)
	fmt.Println(testSlice)
}