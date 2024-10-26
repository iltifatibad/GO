package main

import "fmt"

func main() {

	datas := [2]string{" Hello ", " World "}

	for i := 0; i < len(datas); i++ {
		fmt.Println(i+1, datas[i])
	}
}
