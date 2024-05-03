package main

import "testing"

func TestArea(t *testing.T) {
	inputs := []struct {
		Name    string
		Entrada float64
		Want    float64
	}{
		{
			Name:    "Area 2.00",
			Entrada: 2.00,
			Want:    12.5664,
		},
		{
			Name:    "Area 100.64",
			Entrada: 100.64,
			Want:    31819.3103,
		},
		{
			Name:    "Area 150.00",
			Entrada: 150.00,
			Want:    70685.7750,
		},
	}

	for _, input := range inputs {
		t.Run(input.Name, func(t *testing.T) {
			result := Area(input.Entrada)
			if result != input.Want {
				t.Errorf("want: '%f', got: '%f'", input.Want, result)
			}
		})
	}
}
