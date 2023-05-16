package main

import "fmt"

func main() {
	var a, b [3]int
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])
	var alice, bob int
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alice++
		} else if b[i] > a[i] {
			bob++
		}
	}
	fmt.Println(alice, bob)
}
