package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := (40.0)

	if got != want {
		t.Errorf("Expected %.2f want but received %.2f", want, got)
	}
}



func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{10.0, 10.0},
			want: 100.0,
		},
		{
			name: "Circle",
			shape: Circle{10.0},
			want: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{12.0, 6.0},
			want : 36.0,
		},
	}

	for _, testCase := range areaTests {
		got := testCase.shape.Area()
		expected := testCase.want

		if got != expected {
			t.Errorf("got %g want %g", got, testCase.want)
		}
	}
}