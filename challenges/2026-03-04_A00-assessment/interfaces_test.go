package main

import (
	"math"
	"testing"
)

// #1 Static interface check to catch issue at compile time
// ensures that all shapes satisfy interface and checks at compile time
var _ Shape = Circle{}
var _ Shape = Triangle{}
var _ Shape = Rectangle{}

func TestShapeCaluclations(t *testing.T) {
	tests := []struct {
		name       string
		shape      Shape
		area_want  float64
		perim_want float64
	}{
		{
			name:       "Circle r=1",
			shape:      Circle{r: 1},
			area_want:  math.Pi,
			perim_want: 2 * math.Pi,
		},
		{
			name:       "Rectangle 2x3",
			shape:      Rectangle{a: 2, b: 3},
			area_want:  6,
			perim_want: 10,
		},
		{
			name:       "Triangle 3-4-5",
			shape:      Triangle{a: 3, b: 4, c: 5},
			area_want:  6,
			perim_want: 12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.area_want != tc.shape.Area() {
				t.Errorf("Area calculation ~%.2f for %s is not correct, want ~%.2f", tc.shape.Area(), tc.name, tc.area_want)
			}
			if tc.perim_want != tc.shape.Perimeter() {
				t.Errorf("Perimeter calculation ~%.2f for %s is not correct, want ~%.2f", tc.shape.Perimeter(), tc.name, tc.perim_want)
			}
		})
	}
}

func TestShapeEdgeCases(t *testing.T) {
	t.Run("r0 Edge case", func(t *testing.T) {
		c := Circle{r: 0}
		if c.Area() != 0 || c.Perimeter() != 0 {
			t.Errorf("Expect Area 0 and Perimeter =0 for r=0 Circle")
		}
	})

	t.Run("Negative Shape", func(t *testing.T) {
		r := Rectangle{a: -1, b: 2}
		if r.Area() < 0 || r.Perimeter() < 0 {
			t.Errorf("Phisics broken, negative Area or Perimeter")
		}
	})
}

func TestLargestAreaShape(t *testing.T) {

	t.Run("Empty shape list", func(t *testing.T) {
		if LargestAreaShape([]Shape{}) != nil {
			t.Errorf("Expect nil for empty shape list")
		}
	})

	t.Run("Largest Area Selection", func(t *testing.T) {

		want := Rectangle{10, 20}
		shapes := []Shape{
			want,
			Rectangle{a: 4, b: 3},
			Rectangle{a: 1, b: 2},
		}

		got := LargestAreaShape(shapes)

		if want != got {
			t.Errorf("LargestAreaShape() = %v, want %v ", got, want)
		}

	})
}
