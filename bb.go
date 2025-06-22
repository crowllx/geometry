package geometry

type BB struct {
	L, T, R, B float32
}

func (bb BB) Bounds() (float32, float32) {
	return bb.R - bb.L, bb.B - bb.T
}

func (bb *BB) Translate(v Vector) {
	bb.L += v.X
	bb.R += v.X
	bb.T += v.Y
	bb.B += v.Y
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
