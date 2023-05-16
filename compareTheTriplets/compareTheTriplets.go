package main

import "fmt"

func main() {
	var alice, bob int
	var a, b [3]int
	fmt.Scan(&a[0], &a[1], &a[2], &b[0], &b[1], &b[2])
	for i := range a {
		if a[i] > b[i] {
			alice++
		} else if b[i] > a[i] {
			bob++
		}
	}
	fmt.Println(alice, bob)
}
