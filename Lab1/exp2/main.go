
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Simple Calculator ---")
		fmt.Println("1. Integer Calculator")
		fmt.Println("2. Floating-Point Calculator")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid choice! Enter 1, 2, or 3.")
			continue
		}

		if choice == 3 {
			fmt.Println("Calculator closed.")
			break
		}

		if choice == 1 {
			var a, b int64

			fmt.Print("Enter first integer: ")
			input, _ = reader.ReadString('\n')
			a, err = strconv.ParseInt(strings.TrimSpace(input), 10, 64)
			if err != nil {
				fmt.Println("Invalid integer!")
				continue
			}

			fmt.Print("Enter second integer: ")
			input, _ = reader.ReadString('\n')
			b, err = strconv.ParseInt(strings.TrimSpace(input), 10, 64)
			if err != nil {
				fmt.Println("Invalid integer!")
				continue
			}

			fmt.Println("Addition:", a+b)
			fmt.Println("Subtraction:", a-b)

			if b != 0 && (a > math.MaxInt64/b ||
				a < math.MinInt64/b) {
				fmt.Println("Multiplication: result out of range")
			} else if a == math.MinInt64 && b == -1 ||
				b == math.MinInt64 && a == -1 {
				fmt.Println("Multiplication: result out of range")
			} else {
				fmt.Println("Multiplication:", a*b)
			}

		} else {
			var a, b float64

			fmt.Print("Enter first number: ")
			input, _ = reader.ReadString('\n')
			a, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
			if err != nil || math.IsInf(a, 0) || math.IsNaN(a) {
				fmt.Println("Invalid floating-point number!")
				continue
			}

			fmt.Print("Enter second number: ")
			input, _ = reader.ReadString('\n')
			b, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
			if err != nil || math.IsInf(b, 0) || math.IsNaN(b) {
				fmt.Println("Invalid floating-point number!")
				continue
			}

			fmt.Println("Addition:", a+b)
			fmt.Println("Subtraction:", a-b)
			fmt.Println("Multiplication:", a*b)
		}
	}
}
