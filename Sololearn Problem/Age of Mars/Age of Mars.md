# Age of Mars
> **Sumber :**
> 
> - [Sololearn: Learn to Code](https://www.sololearn.com/)
> - [Go Documentation](https://go.dev/doc/)


## Question
How old would you be on Mars?
A year on Earth has 365 days, while a year on Mars has 687 days.

Create a program that takes your age in Earth years as input, and outputs the corresponding age on Mars.

The given program takes an integer as input and passes it to the `mars_age()` function as argument.
Complete the function to calculate and return the corresponding Mars age.

## Answer
```go
package main

import "fmt"

func MarsAge(x int) int {
	return x * 365 / 687
}

func main() {
	var age int
	fmt.Println("Umur di Bumi")
	fmt.Scanln(&age)

	mars := MarsAge(age)
	fmt.Println("Umur di Bulan")
	fmt.Println(mars)
}
```

**Output:**
```
Umur di Bumi
21
Umur di Bulan
11
```

**Answer step by step:**

Diketahui bahwa dalam 1 tahun, di Bumi terdiri dari 365 hari dan sedangkan di Bulan terdiri dari 687 hari. Maka, logika untuk mengkonversi umur di Bumi ke dalam umur di Bulan adalah $\frac{umurdibumi \times 687}{365}$.
