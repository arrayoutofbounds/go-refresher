// arrays have to be fixed length

package main

import "fmt"

func main() {
	fruitArr2 := [2]string{"Apple", "orange"}

	var fruitArr [2]string

	fruitArr[0] = "apple"
	fruitArr[1] = "orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr2)

	fruitSlice := []string{"Apple", "orange", "grape"}

	fmt.Println(fruitSlice)
}
