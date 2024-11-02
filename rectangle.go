package geometry

type Rect struct {
	Min   Vector
	Max   Vector
}

func (r *Rect) Bounds() (float64, float64) {
	dx := r.Max.X - r.Min.X
	dy := r.Max.Y - r.Min.Y
	return dx, dy
}
func (r *Rect) BB() BB {
	return BB{
		r.Min.X, r.Min.Y, r.Max.X, r.Max.Y,
	}
}

func NewRect(x0, y0, x1, y1 float64) *Rect {
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

func (r *Rect) Collides(s Shape) bool {
    return r.BB().Collides(s)
}
