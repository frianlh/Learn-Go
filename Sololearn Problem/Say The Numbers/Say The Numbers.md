# Say The Numbers
> **Sumber :**
> 
> - [Sololearn: Learn to Code](https://www.sololearn.com/)
> - [Go Documentation](https://go.dev/doc/)


## Question
You are making a robot that can speak numbers. Your robot should take 3 numbers in the range of **0 - 10** as input and output the corresponding texts in English.

## Answer
```go
package main

import "fmt"

func main() {
	var input int
	var i int

	for i = 0; i < 3; i++ {
		// Returns the number of total items successfully scanned and an error if occurred during the read operation.
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

```

**Output:**
```
Input
2
Ouput
Two

Input
5
Output
Five

Input
1
Output
One
```

**Answer step by step:**

Diketahui robot yang dibuat akan meminta 3 input berupa bilangan integer antara 0 - 10. Maka, untuk mentransformasi dari bilangan integer ke dalam teks, cukup dengan menggunakan if statment.
