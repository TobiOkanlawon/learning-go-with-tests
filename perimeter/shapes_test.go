package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("correctly calculates the perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{3.0, 4.0}
		got := rectangle.Perimeter()
		want := 14.0

		if got != want {
			t.Errorf("Expected value %.2f is not equal to %g", want, got)
		}

	})
}

func TestArea(t *testing.T) {

	// refactoring into table tests

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{3.0, 4.0}, hasArea: 12.0},
		{name: "Circle", shape: Circle{7.0}, hasArea: 153.86},
		{name: "Triangle", shape: Triangle{7.0, 8.0}, hasArea: 28.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v expected value %g is not equal to %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}
