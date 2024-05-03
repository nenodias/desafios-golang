package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const PI = 3.14159

func Area(raio float64) float64 {
	result := PI * (raio * raio)
	var saida bytes.Buffer
	_, err := fmt.Fprintf(&saida, "%.4f", result)
	if err != nil {
		panic(err)
	}
	result, err = strconv.ParseFloat(saida.String(), 64)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	var raio float64
	fmt.Scanf("%f", &raio)
	result := Area(raio)
	fmt.Printf("A=%.4f\n", result)
}
