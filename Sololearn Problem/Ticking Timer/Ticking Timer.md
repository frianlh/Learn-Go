# Ticking Timer
> **Sumber :**
> 
> - [Sololearn: Learn to Code](https://www.sololearn.com/)
> - [Go Documentation](https://go.dev/doc/)


## Question
You are building a **Timer** app that should count up to a given number.
Your program needs to take a number as input and make the **Timer** tick that number of times.

The code in main initializes a **Timer** and takes a number as input. Then it calls the `tick()` method for the **Timer** the given number of times.

Define the **Timer** struct with two fields: **id** and **value**, and define the `tick()` method, which should increment the **value** by one and output its current value.

## Answer
```go
package main

import "fmt"

type Timer struct {
	id    string
	value int
}

func (x *Timer) Tick() {
	x.value++            // Increment the value by one
	fmt.Println(x.value) // Output its current value
}

func main() {
	var x int
	fmt.Scanln(&x)

	t := Timer{"timer1", 0}

	for i := 0; i < x; i++ {
		t.Tick()
	}
}
```

**Output:**
```
Input
5

Output
1
2
3
4
5
```
