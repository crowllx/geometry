package geometry

type Circle struct {
	center Vector
	r      float64
	class  ShapeClass
}

func (c Circle) BB() BB {
	return BB{
		c.center.X - c.r,
		c.center.Y - c.r,
		c.center.X + c.r,
		c.center.Y + c.r,
	}
}

func (c Circle) Class() ShapeClass {
	return c.class
}

func (c Circle) Center() Vector {
	return c.center
}

func (c Circle) Radius() float64 {
	return c.r
}

var _ Shape = Circle{}

func NewCircle(c Vector, r float64) Circle {
	return Circle{c, r, CIRCLE}
}
