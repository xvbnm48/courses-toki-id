package main

import "fmt"

func main() {
	var matriks [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&matriks[i][j])
		}
	}

	var transpose [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transpose[i][j] = matriks[j][i] // Perhatikan penggunaan indeks
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d", transpose[i][j])
			if j < 2 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
