package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Shape area is %.2f, and perimeter is %.2f", s.Area(), s.Perimeter())
}

func LargestAreaShape(s []Shape) Shape {
	if len(s) == 0 || s == nil {
		return nil
	}

	s_copy := slices.Clone(s)
	copy(s, s_copy)
	slices.SortFunc(s_copy, func(a, b Shape) int {
		return cmp.Compare(b.Area(), a.Area())
	})
	return s_copy[0]
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
