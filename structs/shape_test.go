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
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{7.0, 13.0, 5.0, 12.0}, 30.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
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
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 9.0}, 42.0},
		{Circle{19.25}, 120.95131716320704},
		{Triangle{7.0, 13.0, 5.0, 12.0}, 25.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
