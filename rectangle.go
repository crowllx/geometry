package geometry

type Rect struct {
	Min Vector
	Max Vector
}

func (r *Rect) Bounds() (float32, float32) {
	dx := r.Max.X - r.Min.X
	dy := r.Max.Y - r.Min.Y
	return dx, dy
}

func (r *Rect) BB() BB {
	return BB{
		r.Min.X, r.Min.Y, r.Max.X, r.Max.Y,
	}
}

func NewRect(x0, y0, x1, y1 float32) *Rect {
	return &Rect{
		NewVector(x0, y0),
		NewVector(x1, y1),
	}
}

var _ Shape = &Rect{}

func (r *Rect) Translate(v Vector) {
	r.Min = r.Min.Add(v)
	r.Max = r.Max.Add(v)
}

func (r *Rect) Rotate(angle float32, origin Vector) {
	r.Min = r.Min.Rotate(angle, origin)
	r.Max = r.Max.Rotate(angle, origin)
}

func (r Rect) Centroid() Vector {
	x := (r.Min.X + r.Max.X) / 2
	y := (r.Min.Y + r.Max.Y) / 2
	return Vector{x, y}
}

func (r *Rect) Collides(s Shape) bool {
	if r.BB().Contains(s.BB()) {
		switch s := s.(type) {
		case *Circle:
			return CircleRectCollision(s, r)
		default:
			return true
		}
	}
	return false
}
