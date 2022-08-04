package main

import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	var lastmatch string
	fmt.Scanln(&lastmatch)

	results = append(results, lastmatch)

	points := 0
	for i := range results {
		switch results[i] {
		case "w":
			points += 3
		case "d":
			points += 1
		case "l":
			points += 0
		}
	}
	fmt.Println(points)
}
