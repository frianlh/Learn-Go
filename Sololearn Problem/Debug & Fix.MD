# Debug & Fix
> **Sumber:**
> 
> - [Sololearn: Learn to Code](https://www.sololearn.com/)
> - [Go Documentation](https://go.dev/doc/)


## Question
Your colleague tried to make a program that is intended to output "**GO**" three times on separate lines. However, he is still new to Go and has made a number of mistakes in the code. Debug and fix the code to produce the desired output.

```go
pack main
import "main"

func mine() {
   // outputs GO 3 times
    fmt.Printline("go")
    fmt.println("go")
    
    //fmn.Println("go")
}
```

## Answer
```go
package main

import "fmt"

func main() {
	// Print GO 3 times
	fmt.Println("GO")
	fmt.Println("GO")
	fmt.Println("GO")
}
```

**Output:**
```
GO
GO
GO
```

**Answer step by step:**
| Sebelum | Sesudah | Keterangan |
|:---|:---|:---|
| `pack main` | `package main` | Gunakan kata kunci `package` |
| `import "main"` | `import "fmt"` | Import package yang salah |
| `func mine()` | `func main()` | `main` package membutuhkan fungsi `main ()` |
| `fmt.Printline("go")` | `fmt.Println("GO")` | Fungsi yang salah |
| `fmt.println("go")` | `fmt.Println("GO")` | go &rarr; GO |
| `//fmn.Println("go")` | `fmt.Println("GO")` | Mengapus comment dan go &rarr; GO|
