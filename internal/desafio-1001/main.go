package main

import "fmt"

func main() {
	var A, B int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	X := Soma(A, B)
	fmt.Printf("X = %d\n", X)
}

func Soma(A, B int) int {
	return A + B
}
