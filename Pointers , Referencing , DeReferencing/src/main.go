package main

import "fmt"

func main() {
	a := 5 // This is my variable and this variable has a location point on memory like as 0xc00005f

	fmt.Println(&a) // If i use &a it gives me location of a on which locate on memory cd ----> 0xc000000a0c8 ( It's referenced data)

	// If i want DeReference this value i should use * for it for example

	fmt.Println(*(&a)) // ----> 5

}
