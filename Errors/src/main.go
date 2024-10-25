package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	result, _ := Even(10)
	if reflect.TypeOf(result).String() != "int" {
		fmt.Println(" Pls Insert Integrer ")
	}

	fmt.Println(" Your Number Is ", result)
}

func Even(x int) (int, error) {
	if x%2 != 0 {
		return 0, errors.New(" There Is A Some Problem ")
	}
	return x, nil
}
