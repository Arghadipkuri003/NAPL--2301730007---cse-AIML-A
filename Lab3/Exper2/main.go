
package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) ReadData() {
	fmt.Print("Enter Name: ")
	fmt.Scan(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)
}

func (p Person) DisplayData() {
	fmt.Println("\n--- Person Details ---")
	fmt.Println("Name   :", p.Name)
	fmt.Println("Age    :", p.Age)
	fmt.Println("Job    :", p.Job)
	fmt.Printf("Salary : %.2f\n", p.Salary)
}

func main() {
	var p1, p2 Person

	fmt.Println("Enter details for Person 1:")
	p1.ReadData()
	p1.DisplayData()

	fmt.Println("\nEnter details for Person 2:")
	p2.ReadData()
	p2.DisplayData()
}
