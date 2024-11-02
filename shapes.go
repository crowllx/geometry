package geometry

type Shape interface {
	Class() ShapeClass
	BB() BB
}

type ShapeClass int

type BB struct {
	L, T, R, B float64
}

func (bb BB) Contains(other BB) bool {
	if bb.R < other.L || bb.L > other.R {
		return false
	}
	if bb.B < other.T || bb.T > other.B {
		return false
	}
	return true
}

const (
	RECT = iota
	CIRCLE
)

type Rect struct {
	Min   Vector
	Max   Vector
	class ShapeClass
}

func (r Rect) Bounds() (float64, float64) {
	dx := r.Max.X - r.Min.X
	dy := r.Max.Y - r.Min.Y
	return dx, dy
}
func (r Rect) BB() BB {
	return BB{
		r.Min.X, r.Min.Y, r.Max.X, r.Max.Y,
	}
}

func (r Rect) Scale(v Vector) Rect {
	r.Min = r.Min.Add(v)
	r.Max = r.Max.Add(v)
	return r
}

func NewRect(x0, y0, x1, y1 float64) Rect {
	return Rect{
		NewVector(x0, y0),
		NewVector(x1, y1),
		RECT,
	}
}

var _ Shape = Rect{}

func (r Rect) Class() ShapeClass {
	return r.class
}
