package main

func main() {
	a, b := calc(100, 80)
	println(a, b)
}

func calc(x, y int) (a, b int) {
	if x%y == 0 {
		return a, b
	}

	a = x / y
	b = x % y
	return a, b

}
