package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)

	delete(emails, "Bob")

	fmt.Println(emails)

	// declare inline
	emails2 := map[string]string{"Bob": "bob@mail.com"}

	fmt.Println(emails2)
}
