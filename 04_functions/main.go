package main

import "fmt"

func greeting(name string) string {
	return "Hello " + "World"
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("ad"))
	fmt.Println(getSum(2, 3))
}
