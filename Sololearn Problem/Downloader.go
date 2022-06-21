package main

import "fmt"

func Download(size int, ch chan int) {
	counts := 0
	for i := 0; i <= size; i++ {
		counts += i
	}
	ch <- counts
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var s1, s2, s3 int
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)

	go Download(s1, ch1)
	go Download(s2, ch2)
	go Download(s3, ch3)

	fmt.Println(<-ch1 + <-ch2 + <-ch3)
}
