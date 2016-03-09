package rectanglePkg

// capital letter - func or var is visable from outside the package
// small letter - func or var is not visble
type rectangle struct {
  x int
  y int
}

func (r rectangle) Area() int {
  return r.x*r.y
}

func (r rectangle) Circ() int {
  return 2*r.x + 2*r.y
}

// This function is not visable
func newRect (x, y int) rectangle {
  return rectangle{x:x, y:y}
}

// This function is visable
func NewRectangle (x, y int) rectangle {
    return newRect(x, y)
}
