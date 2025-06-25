package main

import "fmt"

func main() {
	var name = "Rajeev Singh" 
	fmt.Printf("Variable 'name' is of type %T\n", name)

	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
		firstName, lastName, age, salary)
}
