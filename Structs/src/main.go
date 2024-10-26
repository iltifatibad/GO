package main

import "fmt"

func main() {
	var employee struct {
		name      string
		age       int
		isMarried bool
	}

	employee.name = " Hello World "
	employee.age = 20
	employee.isMarried = true

	// Or Same Time We Can Crete Our Own Type For Example As String

	type employee_type struct {
		name      string
		age       int
		isMarried bool
	}

	var employee1 employee_type

	employee1.name = " Hello "
	employee1.age = 20
	employee1.isMarried = true

	// Or We Can Use Slices In Struct And Type Structs For Example

	type employee_type_slices struct {
		name      string
		age       int
		isMarried bool
		kids      []string
	}

	employee2 := employee_type_slices{
		name:      " Hello ",
		age:       5,
		isMarried: true,
		kids: []string{
			" Hello ",
			" World ",
		},
	}

	fmt.Println(employee2.age)
}
