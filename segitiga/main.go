package main

import "fmt"

func main() {
	var a, t int

	fmt.Scan(&a, &t)
	luas := float32(a*t) / 2
	fmt.Printf("%.2f\n", luas)
}
