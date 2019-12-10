package main

import (
	"fmt"
	"math"

	"github.com/arrayoutofbounds/go_crash_course/03_packages/strutil"
)

// it runs this function when you run the file
func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(strutil.Reverse("hello"))
}
