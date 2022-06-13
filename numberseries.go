package main

import "fmt"

func main() {
	a := 1
	b := 1
	fmt.Println("Enter the number")
	var number int
	fmt.Scanln(&number)
	fmt.Println(a)
	fmt.Println(b)

	for i := 3; i <= number; i++ {

		pattern := a + b

		fmt.Println(pattern)
		a = b
		b = pattern

	}

}
