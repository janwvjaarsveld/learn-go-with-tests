package structs

import "testing"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type ShapeTest struct {
	shape Shape
	name  string
	want  float64
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []ShapeTest{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 40},
		{name: "Circle", shape: Circle{100}, want: 628.3185307179587},
	}
	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []ShapeTest{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
