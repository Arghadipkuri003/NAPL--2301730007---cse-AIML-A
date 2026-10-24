package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan string, 3)

	// Goroutine 1: Calculating squares
	go func() {
		fmt.Println("Squares goroutine started")

		var squares []int
		for i := 1; i <= 5; i++ {
			squares = append(squares, i*i)
		}

		time.Sleep(300 * time.Millisecond)
		fmt.Println("Squares goroutine completed")
		results <- fmt.Sprintf("Squares: %v", squares)
	}()

	// Goroutine 2: Calculating cubes
	go func() {
		fmt.Println("Cubes goroutine started")

		var cubes []int
		for i := 1; i <= 5; i++ {
			cubes = append(cubes, i*i*i)
		}

		time.Sleep(100 * time.Millisecond)
		fmt.Println("Cubes goroutine completed")
		results <- fmt.Sprintf("Cubes: %v", cubes)
	}()

	// Goroutine 3: Fibonacci numbers
	go func() {
		fmt.Println("Fibonacci goroutine started")

		fib := []int{0, 1}
		for i := 2; i < 7; i++ {
			fib = append(fib, fib[i-1]+fib[i-2])
		}

		time.Sleep(200 * time.Millisecond)
		fmt.Println("Fibonacci goroutine completed")
		results <- fmt.Sprintf("Fibonacci: %v", fib)
	}()

	fmt.Println("\nResults received in main():")
	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}

	close(results)
}
