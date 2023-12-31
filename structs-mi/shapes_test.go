package structsmi

import "testing"

func TestPerimeter(t *testing.T) {
  rectangle := Rectangle{10.0, 10.0}
  got := Perimeter(rectangle)

  expected := 40.0

  if got != expected {
    t.Errorf("expected %.2f but got %.2f", expected, got)
  }

}

func TestArea(t *testing.T) {
  areaTest := []struct{
    name string
    shape Shape
    want float64
  } {
    {name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
    {name: "Circle", shape: Circle{10}, want: 314.1592653589793},
    {name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
  }

  for _, tt := range areaTest{
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Area()
      if got != tt.want {
			  t.Errorf("got %g want %g", got, tt.want)
		  }
    })
  }
}

/* func TestArea(t *testing.T) {

  checkArea := func (t testing.TB, shape Shape, want float64)  {
    t.Helper()

    got := shape.Area()
    if got != want {
			t.Errorf("got %g want %g", got, want)
		}
  }

  t.Run("rectangles", func(t *testing.T) {    
    rectangle := Rectangle{12, 6}
    checkArea(t, rectangle, 72.0)
  })

  t.Run("circles", func(t *testing.T) {
    circle := Circle{10}
    checkArea(t, circle, 314.1592653589793)
  })
} */
