
package main

import "fmt"

func modifyValue(ptr *int) {
	*ptr = 100
}

type Student struct {
	ID   int
	Age  int
	Name string
}

func main() {
	num := 10
	ptr := &num

	fmt.Println("Original value:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Pointer value (address):", ptr)
	fmt.Println("Dereferenced value:", *ptr)

	fmt.Println("\nBefore modification:", num)
	modifyValue(ptr)
	fmt.Println("After modification:", num)

	studentPtr := new(Student)

	studentPtr.ID = 101
	studentPtr.Age = 21
	studentPtr.Name = "Rahul"

	fmt.Println("\nStudent Details:")
	fmt.Println("ID:", studentPtr.ID)
	fmt.Println("Age:", studentPtr.Age)
	fmt.Println("Name:", studentPtr.Name)

	studentPtr.Age = 22
	studentPtr.Name = "Aman"

	fmt.Println("\nUpdated Student Details:")
	fmt.Println("Age:", studentPtr.Age)
	fmt.Println("Name:", studentPtr.Name)
}
