# Match Results
> **Sumber :**
> 
> - [Sololearn: Learn to Code](https://www.sololearn.com/)
> - [Go Documentation](https://go.dev/doc/)


## Question
You are making a program to analyze sport match results and calculate the points of the given team.
The match results are stored in an array called **results**.
Each match has one of the following results:
- "w" - won
- "l" - lost
- "d" - draw

A win adds three points, a draw adds one point, and a lost match does not add any points.

Your program needs to take the last match result as input and append it to the results array. After that, calculate and output the total points the team gained from the results.e.

## Answer
```go
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
```

**Output:**
```
Input
w

Output
22
```

**Answer step by step:**

Diketahui bahwa nilai dari hasil sport match mengikuti aturan sebagai berikut "w" - won: 3 poin, "l" - lost: 0 poin, dan "d" - draw: 1 poin. Maka, logika untuk menyelesaikan masalah ini adalah melakukan looping untuk mengecek hasil dari masing-masing sport match (termasuk match terakhir/input) untuk mengetahui poin pada match tersebut, lalu store poin tersebut pada suatu variabel untuk mengetahui total poin yang didapatkan.
