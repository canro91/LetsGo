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
	t.Run("Rectangle_ReturnsArea", func(t *testing.T){
		expected := 12.0
		rectangle := Rectangle { 3.0, 4.0 }
		area := rectangle.Area()
		
		if expected != area {
			t.Errorf("Expected %.2f but was %.2f", expected, area)
		}
	})

	t.Run("Circle_ReturnsArea", func(t *testing.T){
		expected := math.Pi
		circle := Circle { 1.0 }
		area := circle.Area()
		
		if expected != area {
			t.Errorf("Expected %.2f but was %.2f", expected, area)
		}
	})
}