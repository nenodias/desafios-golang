package main

import "testing"

func TestSoma(t *testing.T) {
	inputs := []struct {
		Name string
		A    int
		B    int
		Want int
	}{
		{
			Name: "10 + 9",
			A:    10,
			B:    9,
			Want: 19,
		},
		{
			Name: "-10 + 4",
			A:    -10,
			B:    4,
			Want: -6,
		},
		{
			Name: "15 - 7",
			A:    15,
			B:    -7,
			Want: 8,
		},
	}

	for _, input := range inputs {
		t.Run(input.Name, func(t *testing.T) {
			result := Soma(input.A, input.B)
			if result != input.Want {
				t.Errorf("want: '%d', got: '%d'", input.Want, result)
			}
		})
	}
}
