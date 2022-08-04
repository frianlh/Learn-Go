package main

import "fmt"

func MarsAge(x int) int {
	return x * 365 / 687
}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := MarsAge(age)
	fmt.Println(mars)
}
