package main

import (
	"fmt"
)

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	banyakBebek := n / m
	sisa := n % m

	fmt.Printf("masing-masing %d\n", banyakBebek)
	fmt.Printf("bersisa %d\n", sisa)
}
