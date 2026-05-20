package structs

import "testing"

// Testing Area
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{9.0, 4.5}
		want := 40.5
		checkArea(t, rectangle, want)
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{19.25}
		want := 1164.1564276958677
		checkArea(t, circle, want)
	})

	// Table Driven Tests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{SideA: 7.0, SideB: 13.0, Base: 5.0, Height: 12.0}, hasArea: 30.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}

// Testing Perimeter
func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangle Perimeter", func(t *testing.T) {
		rectangle := Rectangle{9.0, 4.5}
		want := 27.0
		checkPerimeter(t, rectangle, want)
	})

	t.Run("Circle Perimeter", func(t *testing.T) {
		circle := Circle{19.25}
		want := 120.95131716320704
		checkPerimeter(t, circle, want)
	})

	// Table Driven Test
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 9.0}, hasPerimeter: 42.0},
		{name: "Circle", shape: Circle{Radius: 19.25}, hasPerimeter: 120.95131716320704},
		{name: "Triangle", shape: Triangle{SideA: 7.0, SideB: 13.0, Base: 5.0, Height: 12.0}, hasPerimeter: 25.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("got %g want %g", got, tt.hasPerimeter)
		}
	}
}
