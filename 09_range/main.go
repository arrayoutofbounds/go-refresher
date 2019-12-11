package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6}

	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	// add ids
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum)

	// range with map
	emails2 := map[string]string{"Bob": "bob@mail.com"}
	for k, v := range emails2 {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails2 {
		fmt.Println("Name " + k)
	}
}
