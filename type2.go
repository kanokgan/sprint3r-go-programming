package main

import "fmt"

type Contact struct {
	Name string
	Email string
}

func FullContact(contact Contact) string {
	return contact.Name + " " + contact.Email
}

func main() {
	var contact = Contact {					// value type
		Name: "Kanokgan",
		Email: "admin@kanokgan.com",
	}

	fmt.Println(contact)


	var contact2 = &Contact {				// reference type
		Name: "Kanokgan",
		Email: "admin@kanokgan.com",
	}

	fmt.Println(contact2)

	var contact3 = contact2
	contact3.Name = "xxx"

	fmt.Println(contact)
	fmt.Println(contact2)
	fmt.Println(contact3)

	fmt.Println(FullContact(contact))
}




