package geometry

type Circle struct {
	center Vector
	r      float64
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

func (c Circle) Radius() float64 {
	return c.r
}

var _ Shape = &Circle{}

func NewCircle(c Vector, r float64) *Circle {
	return &Circle{c, r}
}

func (c *Circle) Translate(v Vector) {
	c.center = c.center.Add(v)
}

func (c *Circle) Collides(s Shape) bool {
	if c.BB().Contains(s.BB()) {
		switch s.(type) {
		case *Rect:
			return CircleRectCollision(c, s.(*Rect))
		case *BB:
			return CircleBBCollision(c, s.(*BB))
		case *Circle:
			return CircleCircleCollision(c, s.(*Circle))
		default:
			return false
		}
	}
    return false

}
