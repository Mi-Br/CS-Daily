package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

// ---

// ### Exercise 3: Interface + Multiple Types
// Define a `Shape` interface with `Area() float64` and `Perimeter() float64`.

// Implement it for:
// - `Circle`
// - `Rectangle`
// - `Triangle` (three sides, use Heron's formula for area)

// Write a function `PrintShapeInfo(s Shape)` that prints both values.

// Then write a function `LargestArea(shapes []Shape) Shape` that returns the shape with the largest area.

// ---

type Shape interface {
	Area() float64
	Perimeter() float64
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Shape area is %.2f, and perimeter is %.2f", s.Area(), s.Perimeter())
}

func LargestAreaShape(slist []Shape) Shape {
	// if len(slist) == 0 {
	// 	return nil
	// }
	// largest_index := 0
	// for i, _ := range slist {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if slist[i].Area() > slist[largest_index].Area() {
	// 		largest_index = i
	// 	}
	// }
	// return slist[largest_index]

	//maybe this one is better, more compact
	slices.SortFunc(slist, func(a, b Shape) int {
		return cmp.Compare(a.Area(), b.Area())
	})
	return slist[0]
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

type Rectangle struct {
	a, b float64
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.a + r.b)
}
