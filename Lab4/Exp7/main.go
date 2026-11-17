package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero: denominator cannot be zero")
	}
	return a / b, nil
}

func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative square root: cannot calculate the square root of %g", n)
	}
	return math.Sqrt(n), nil
}

func ParseInput(input string) (float64, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return 0, fmt.Errorf("empty input: please enter a number")
	}

	value, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input %q: %w", input, err)
	}
	return value, nil
}

func main() {
	fmt.Println("Experiment 7: Error Handling in Go")

	if result, err := Divide(10, 2); err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}
	if _, err := Divide(10, 0); err != nil {
		fmt.Println("Expected error:", err)
	}

	if result, err := SquareRoot(25); err != nil {
		fmt.Println("Square-root error:", err)
	} else {
		fmt.Println("Square root of 25 =", result)
	}
	if _, err := SquareRoot(-4); err != nil {
		fmt.Println("Expected error:", err)
	}

	if result, err := ParseInput("42"); err != nil {
		fmt.Println("Input error:", err)
	} else {
		fmt.Println("Parsed input =", result)
	}
	if _, err := ParseInput("   "); err != nil {
		fmt.Println("Expected error:", err)
	}
}
