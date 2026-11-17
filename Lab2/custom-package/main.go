
package main

import (
	"fmt"
	"custom-package/myutils"
)

func main() {
	text := "Hello World"

	fmt.Println("Original String:", text)
	fmt.Println("Reversed String:", myutils.Reverse(text))
	fmt.Println("Number of Vowels:", myutils.CountVowels(text))

	fact, err := myutils.Factorial(5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial of 5:", fact)
	}

	power, err := myutils.Power(2, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("2 raised to 3:", power)
	}
}
