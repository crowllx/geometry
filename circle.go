package geometry

type Circle struct {
	center Vector
	r      float32
}

func (c Circle) BB() BB {
	return BB{
		c.center.X - c.r,
		c.center.Y - c.r,
		c.center.X + c.r,
		c.center.Y + c.r,
	}
}

func (c Circle) Center() Vector {
	return c.center
}

func (c Circle) Radius() float32 {
	return c.r
}

func (c *Circle) Rotate(angle float32, origin Vector) {
	c.center = c.center.Rotate(angle, origin)
}

var _ Shape = &Circle{}

func NewCircle(c Vector, r float32) *Circle {
	return &Circle{c, r}
}

func (c *Circle) Translate(v Vector) {
	c.center = c.center.Add(v)
}

func (c *Circle) Collides(s Shape) bool {
	if c.BB().Contains(s.BB()) {
		switch s := s.(type) {
		case *Rect:
			return CircleRectCollision(c, s)
		case *Circle:
			return CircleCircleCollision(c, s)
		default:
			return false
		}
	}
	return false

}
