package structsMethodsInterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	shapes := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 20.0}, 60},
		{Circle{10.0}, 62.83185307179586},
	}

	for _, tt := range shapes {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got: %.2f, want: %.2f", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	shapes := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{20.0, 10.0}, want: 200.0},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
	}

	for _, tt := range shapes {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v, got: %.2f, want: %.2f", tt.shape, got, tt.want)
		}
	}

}
