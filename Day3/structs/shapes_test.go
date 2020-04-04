package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T){
	expected := 40.0
	rectangle := Rectangle{ 10.0, 10.0 }
	perimeter := Perimeter(rectangle)

	if expected != perimeter {
		t.Errorf("Expected %.2f but was %.2f", expected, perimeter)
	}
}

func TestArea(t *testing.T){
	assertArea := func(t *testing.T, shape Shape, expectedArea float64){
		t.Helper()

		area := shape.Area()
		
		if expectedArea != area {
			t.Errorf("%#v Expected %.2f but was %.2f", shape, expectedArea, area)
		}
	}

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{ name: "Rectangle", shape: Rectangle { 3.0, 4.0 }, hasArea: 12.0 },
		{ name: "Circle",    shape: Circle { 1.0 },         hasArea: math.Pi },
		{ name: "Triangle",  shape: Triangle{3.0, 4.0},     hasArea: 6.0 },
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			assertArea(t, tt.shape, tt.hasArea)	
		})
	}
}