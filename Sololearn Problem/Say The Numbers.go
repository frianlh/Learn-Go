package main

import "fmt"

func main() {
	var input int
	var i int

	for i = 0; i < 3; i++ {
		n, err := fmt.Scanln(&input)
		if n == 0 || err != nil {
			break
		}
		switch input {
		case 0:
			fmt.Print("Zero")
		case 1:
			fmt.Print("One")
		case 2:
			fmt.Print("Two")
		case 3:
			fmt.Print("Three")
		case 4:
			fmt.Print("Four")
		case 5:
			fmt.Print("Five")
		case 6:
			fmt.Print("Six")
		case 7:
			fmt.Print("Seven")
		case 8:
			fmt.Print("Eight")
		case 9:
			fmt.Print("Nine")
		case 10:
			fmt.Print("Ten")
		default:
		}
		if i < 2 {
			fmt.Println("")
		}
	}
}
