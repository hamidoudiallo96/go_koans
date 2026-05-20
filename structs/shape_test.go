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
}
