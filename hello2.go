package main

import "fmt"

type Contact struct {
	Name string
	Email string
}

func main() {
	var contact = Contact {
		Name: "Kanokgan",
		Email: "admin@kanokgan.com",
	}

	fmt.Println(contact)
}