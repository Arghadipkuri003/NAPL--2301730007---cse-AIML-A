
package main

import "fmt"

func main() {
	// Part 1: Slice operations
	fmt.Println("=== Slice Operations ===")

	students := []string{"Aman", "Rahul", "Priya"}
	fmt.Println("Original slice:", students)

	students = append(students, "Neha")
	fmt.Println("After adding Neha:", students)

	index := 1
	students = append(students[:index], students[index+1:]...)
	fmt.Println("After removing index 1:", students)

	students[1] = "Sneha"
	fmt.Println("After updating index 1:", students)

	// Part 2: Map operations
	fmt.Println("\n=== Map Operations ===")

	marks := map[string]int{
		"Maths":   85,
		"Science": 90,
		"English": 78,
	}
	fmt.Println("Original map:", marks)

	marks["Computer"] = 95
	fmt.Println("After inserting Computer:", marks)

	delete(marks, "English")
	fmt.Println("After deleting English:", marks)

	subject := "Maths"
	if mark, found := marks[subject]; found {
		fmt.Println(subject, "marks:", mark)
	} else {
		fmt.Println(subject, "not found")
	}
}
